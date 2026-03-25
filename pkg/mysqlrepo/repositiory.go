package mysqlrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

const (
	FormatLike = "%%%s%%"
)

var ErrNothingChanged = errors.New("nothing changed")

type SQLStore interface {
	sqlx.Queryer
	sqlx.Execer
}

type TxSQLStore interface {
	sqlx.Queryer
	sqlx.Execer

	Rollback() error
	Commit() error
}

type Repo[T any] struct {
	SQLStore SQLStore
	RCli     *redis.Client
	Logger   Loggerer
	TInfo    *DBTableInfo
	//config
	conf Config
}

type Config struct {
	DebugMode bool
}

func BeginTx(sqlstore SQLStore) (TxSQLStore, error) {
	conn, ok := sqlstore.(*sqlx.DB)
	if !ok {
		return nil, errors.New("SQLStore not a *sqlx.DB")
	}

	tx, err := conn.Beginx()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func Select(store SQLStore, dest any, query string, args ...any) error {
	return sqlx.Select(store, dest, query, args...)
}

func Get(store SQLStore, dest any, query string, args ...any) error {
	return sqlx.Get(store, dest, query, args...)
}

func New[T any](sqlStore SQLStore, table string, config Config, m T) *Repo[T] {
	if sqlStore == nil {
		panic("sqlStore can't be nil")
	}

	if table == "" {
		panic("table can't be empty string")
	}

	return &Repo[T]{
		SQLStore: sqlStore,
		conf:     config,
		TInfo:    PrepareDBTableInfo(table, m),
		Logger:   DefaultLogger(),
	}
}

func (br *Repo[T]) SetDebugMode(debug bool) {
	br.conf.DebugMode = debug
}

func (br *Repo[T]) Clone() *Repo[T] {
	return &Repo[T]{
		SQLStore: br.SQLStore,
		RCli:     br.RCli,
		Logger:   br.Logger,
		TInfo:    br.TInfo,
		conf:     br.conf,
	}
}

func (br *Repo[T]) CloneWithStore(store SQLStore) *Repo[T] {
	return &Repo[T]{
		SQLStore: store,
		RCli:     br.RCli,
		Logger:   br.Logger,
		TInfo:    br.TInfo,
		conf:     br.conf,
	}
}

func (br *Repo[T]) SetLogger(l Loggerer) error {
	br.Logger = l
	return nil
}

func (br *Repo[T]) SetStore(store SQLStore) error {
	br.SQLStore = store
	return nil
}

func (br *Repo[T]) BeginTx() (TxSQLStore, error) {
	return BeginTx(br.SQLStore)
}

func (br *Repo[T]) TxRepo() (*Repo[T], error) {
	tx, err := br.BeginTx()
	if err != nil {
		return nil, err
	}

	txRepo := br.CloneWithStore(tx)
	return txRepo, nil
}

func (br *Repo[T]) Commit() error {
	tx, ok := br.SQLStore.(TxSQLStore)
	if !ok {
		return errors.New("SQLStore not a *sqlx.Tx")
	}

	return tx.Commit()
}

func (br *Repo[T]) Rollback() error {
	tx, ok := br.SQLStore.(TxSQLStore)
	if !ok {
		return errors.New("SQLStore not a *sqlx.Tx")
	}

	return tx.Rollback()
}

func (br *Repo[T]) SetRCli(rCli *redis.Client) error {
	br.RCli = rCli
	return nil
}

func (br *Repo[T]) AppendSoftDelete(cond Conditions) Conditions {
	if br.TInfo.DeleteTime.TimeType != None {
		return append(cond, sq.Eq{fmt.Sprintf("%s.%s", br.TInfo.Table, br.TInfo.DeleteTimeKey()): br.TInfo.DeleteTimeEmpty()})
	}
	return cond
}

func (br *Repo[T]) AppendExtraCondtions(cond Conditions) Conditions {
	return br.AppendSoftDelete(cond)
}

func (br *Repo[T]) SelectClause() SelectBuilder {
	return NewSelectClause().Columns(br.TInfo.Columns...).From(br.TInfo.Table)
}

func (br *Repo[T]) InsertCluase() InsertBuilder {
	return NewInsertClause().Into(br.TInfo.Table)
}

func (br *Repo[T]) UpdateClause() UpdateBuilder {
	return NewUpdateClause().Table(br.TInfo.Table)
}

func (br *Repo[T]) DeleteClause() DeleteBuilder {
	return NewDeleteClause().From(br.TInfo.Table)
}

func (br *Repo[T]) Count(b SelectBuilder, f Lister) (int, error) {
	var total int
	cb := b.RemoveColumns().Column("count(*) as total")

	stm, args, _ := cb.ToSql()

	if strings.Contains(stm, "GROUP BY") {
		stm = fmt.Sprintf(`
			SELECT COUNT(*) FROM (%s) AS x
		`, stm)
	}

	r := br.SQLStore.QueryRowx(stm, args...)
	br.PrintLog(stm, args...)
	if err := r.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}

func (br *Repo[T]) Select(b Builder) ([]T, error) {
	stm, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	br.PrintLog(stm, args...)

	dest := make([]T, 0, 1)
	err = Select(br.SQLStore, &dest, stm, args...)
	return dest, err
}

func ApplyFilter(b SelectBuilder, filters map[string][]string, columnMap map[string]string) SelectBuilder {
	if fl := len(filters); fl > 0 {
		for k, v := range filters {
			table, e := columnMap[k]
			if !e { //filtering something we dont have
				continue
			}
			if table != "" {
				b = b.AddIn(fmt.Sprintf("%s.%s", table, k), v)
			} else {
				b = b.AddIn(k, v)
			}
		}
	}
	return b
}

func (br *Repo[T]) ApplyFilter(b SelectBuilder, filters map[string][]string) SelectBuilder {
	return ApplyFilter(b, filters, br.TInfo.ColumnMap)
}

func ApplyKeyword(b SelectBuilder, keyword string, searchColumns []string) SelectBuilder {
	if keyword != "" {
		if len(searchColumns) > 0 {
			matchExpr := MatchWhere(searchColumns, keyword)
			b = b.Add(matchExpr)
		}
	}
	return b
}

func (br *Repo[T]) ApplyKeyword(b SelectBuilder, keyword string) SelectBuilder {
	return ApplyKeyword(b, keyword, br.TInfo.SearchableColumns)
}

func (br *Repo[T]) ApplyRAGUUIDs(b SelectBuilder, uuids []string) SelectBuilder {
	if len(uuids) > 0 {
		b = b.AddIn("uuid", uuids)
		uuidsStr := make([]string, len(uuids))
		for i, v := range uuids {
			uuidsStr[i] = fmt.Sprintf("'%s'", v)
		}
		b = b.OrderByClause(sq.Expr("FIELD(uuid, " + strings.Join(uuidsStr, ",") + ")"))
	}
	return b
}

func ApplyOrder(b SelectBuilder, sort []string, columnMap map[string]string, defaultOrder string) SelectBuilder {
	if len(sort) > 0 {
		orders := marshalOrders(sort)
		validOrders := []string{}
		for _, o := range orders {
			if _, e := columnMap[strings.Split(o, " ")[0]]; e {
				validOrders = append(validOrders, o)
			}
		}
		b = b.OrderBy(validOrders...)
	} else if defaultOrder != "" {
		// TODO: add default order by
		b = b.OrderBy(defaultOrder)
	}

	return b
}

func (br *Repo[T]) ApplyOrder(b SelectBuilder, sort []string) SelectBuilder {
	var defaultOrder string
	if br.TInfo.CreateTime.TimeType != None {
		defaultOrder = fmt.Sprintf("%s.%s DESC", br.TInfo.Table, br.TInfo.CreateTime.Key)
	}
	return ApplyOrder(b, sort, br.TInfo.ColumnMap, defaultOrder)
}

func (br *Repo[T]) List(b SelectBuilder, f Lister) ([]T, int, error) {
	b = br.ApplyFilter(b, f.Filters())
	// b = br.ApplyKeyword(b, f.Keyword())
	b = br.ApplyRAGUUIDs(b, f.RAGUUIDs())
	b = br.ApplyOrder(b, f.Sort())

	b = b.SetConditions(br.AppendExtraCondtions(b.Conditions()))

	stm, args, _ := b.ToSql()
	fmt.Printf("stm: %s, args: %v\n", stm, args)

	totalRows, err := br.Count(b, f)
	if err != nil {
		return nil, 0, err
	}

	if f.Limit() > 0 {
		b = b.Limit(uint64(f.Limit()))
	}

	if f.Offset() >= 0 && f.Limit() > 0 {
		b = b.Offset(uint64(f.Offset()))
	}

	result, err := br.Select(b)
	return result, totalRows, err
}

func (br *Repo[T]) GetAll(b SelectBuilder) ([]T, error) {
	return br.Select(b.SetConditions(br.AppendExtraCondtions(b.Conditions())))
}

// TODO :: CACHING
func (br *Repo[T]) Get(b SelectBuilder) (dest T, err error) {
	stm, args, err := b.SetConditions(br.AppendExtraCondtions(b.Conditions())).ToSql()
	if err != nil {
		return
	}

	br.PrintLog(stm, args...)

	err = Get(br.SQLStore, &dest, stm, args...)
	return
}

func (br *Repo[T]) Exec(b Builder) (sql.Result, error) {
	stm, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	result, err := br.SQLStore.Exec(stm, args...)
	br.PrintLog(stm, args...)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (br *Repo[T]) Create(data map[string]any, updateFields ...string) (sql.Result, error) {
	br.setDefaultCreateTime(data)
	br.setDefaultUpdateTime(data)
	return br.Exec(br.InsertCluase().SetMap(data).OnDuplicateKeyUpdateSuffix(updateFields...))
}

func (br *Repo[T]) CreateFromModel(m T, updateFields ...string) (sql.Result, error) {
	b := br.InsertCluase().Columns(br.TInfo.InsertColumns...)
	values := make([]any, 0, len(br.TInfo.InsertColumns))
	model_data := ExtractDBTag(m)
	for _, column := range br.TInfo.InsertColumns {
		if _, e := model_data[column]; !e {
			return nil, fmt.Errorf("data for column \"%s\" not found", column)
		}
		values = append(values, model_data[column])
	}

	stm, args, _ := b.Values(values...).OnDuplicateKeyUpdateSuffix(updateFields...).ToSql()
	fmt.Printf("stm: %s, args: %v\n", stm, args)

	return br.Exec(b.Values(values...).OnDuplicateKeyUpdateSuffix(updateFields...))
}

func (br *Repo[T]) BatchCreate(m []map[string]any, updateFields ...string) (sql.Result, error) {
	if len(m) == 0 {
		return nil, nil
	}
	b := br.InsertCluase().Columns(br.TInfo.InsertColumns...)
	for _, data := range m {
		values := make([]any, 0, len(br.TInfo.InsertColumns))

		br.setDefaultCreateTime(data)
		br.setDefaultUpdateTime(data)

		for _, column := range br.TInfo.InsertColumns {
			if _, e := data[column]; !e {
				return nil, fmt.Errorf("data for column \"%s\" not found", column)
			}
			values = append(values, data[column])
		}
		b = b.Values(values...)
	}
	return br.Exec(b.OnDuplicateKeyUpdateSuffix(updateFields...))
}

func (br *Repo[T]) BatchCreateFromModel(list []T, updateFields ...string) (sql.Result, error) {
	dataList := []map[string]any{}
	for i := range list {
		dataList = append(dataList, ExtractDBTag(list[i]))
	}
	return br.BatchCreate(dataList, updateFields...)
}

func (br *Repo[T]) Update(ub UpdateBuilder, data map[string]any) (sql.Result, error) {
	br.setDefaultUpdateTime(data)
	return br.Exec(ub.Table(br.TInfo.Table).SetMap(data).SetConditions(br.AppendExtraCondtions(ub.Conditions())))
}

func (br *Repo[T]) UpdateFromModel(ub UpdateBuilder, m T) (sql.Result, error) {
	return br.Update(ub, ExtractDBTag(m))
}

func (br *Repo[T]) MustUpdate(ub UpdateBuilder, data map[string]any) error {
	result, err := br.Update(ub, data)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return errors.New("nothing changed")
	}

	return nil
}

func (br *Repo[T]) Delete(db DeleteBuilder) (sql.Result, error) {
	return br.Exec(db.From(br.TInfo.Table))
}

func (br *Repo[T]) MustDelete(db DeleteBuilder) error {
	result, err := br.Delete(db)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return ErrNothingChanged
	}

	return nil
}

// TODO:: must use transaction
func (br *Repo[T]) BatchDelete(conds []Conditions) error {
	return errors.New("func is not implemented")
}

// SoftDelete marks the row as deleted instead of actually deleting it
func (br *Repo[T]) SoftDelete(b DeleteBuilder) (sql.Result, error) {
	if br.TInfo.DeleteTime.TimeType == None {
		return nil, errors.New("soft delete not supported")
	}
	return br.Update(NewUpdateClause().Table(br.TInfo.Table).SetConditions(b.Conditions()), map[string]any{br.TInfo.DeleteTime.Key: br.TInfo.DeleteTimeNow()})
}

func (br *Repo[T]) MustSoftDelete(b DeleteBuilder) error {
	if br.TInfo.DeleteTime.TimeType == None {
		return errors.New("soft delete not supported")
	}
	return br.MustUpdate(NewUpdateClause().Table(br.TInfo.Table).SetConditions(b.Conditions()), map[string]any{br.TInfo.DeleteTime.Key: br.TInfo.DeleteTimeNow()})
}

func (br *Repo[T]) PrintLog(stmt string, args ...any) {
	if br.conf.DebugMode && br.Logger != nil {
		br.Logger.Printf(stmt, args...)
	}
}

func (br *Repo[T]) setDefaultCreateTime(data map[string]any) {
	if br.TInfo.CreateTime.TimeType != None {
		if _, e := data[br.TInfo.CreateTime.Key]; !e {
			data[br.TInfo.CreateTime.Key] = br.TInfo.CreateTimeNow()
		}
	}
}

func (br *Repo[T]) setDefaultUpdateTime(data map[string]any) {
	if AutoTime(br.TInfo.UpdateTime.TimeType) {
		if _, e := data[br.TInfo.UpdateTime.Key]; !e {
			data[br.TInfo.UpdateTime.Key] = br.TInfo.UpdateTimeNow()
		}
	}
}
