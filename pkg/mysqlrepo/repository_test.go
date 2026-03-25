package mysqlrepo

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Table1 struct {
	ID          int    `table:"test_table" db:"id" json:"id"`
	Name        string `table:"test_table" db:"name" dbfield:"search" json:"name"`
	Description string `table:"test_table" db:"description" dbfield:"search" json:"description"`
	Table2ID    int    `table:"test_table" db:"table2Id" json:"table2Id"`
	// EntityUnix
}
type Table1List []Table1

func (l *Table1List) Scan(val any) error {
	*l = make(Table1List, 0, 1)
	switch v := val.(type) {
	case string:
		return json.Unmarshal([]byte(v), l)
	case []byte:
		return json.Unmarshal(v, l)
	default:
		fmt.Println(val)
		fmt.Println(v)
	}
	return nil
}

type Table2 struct {
	ID   int    `db:"id"`
	Misc string `db:"misc"`
	// EntityTime
}
type Table2List []Table2

func (l *Table2List) Scan(val any) error {
	*l = make(Table2List, 0, 1)
	switch v := val.(type) {
	case string:
		return json.Unmarshal([]byte(v), l)
	case []byte:
		return json.Unmarshal(v, l)
	default:
		fmt.Println(val)
		fmt.Println(v)
	}
	return nil
}

type TestListForm struct {
	Page         int                 `form:"page"`
	PerPage      int                 `form:"perpage"`
	SearchPhrase string              `form:"keyword"`
	sort         []string            `form:"sort"`
	filters      map[string][]string `form:"filters"`
	embedding    []float32           `form:"embedding"`
	ragUUIDs     []string            `form:"ragUUIDs"`
}

type TestSearchForm struct {
	TestListForm
}

// Sorts implements Lister.
func (tf TestSearchForm) Sort() []string {
	return tf.sort
}
func (tf TestSearchForm) Filters() map[string][]string {
	return tf.filters
}
func (tf TestSearchForm) Keyword() string {
	return tf.SearchPhrase
}
func (tf TestSearchForm) Limit() int {
	return tf.PerPage
}
func (tf TestSearchForm) Offset() int {
	return max(0, tf.Page-1) * tf.Limit()
}
func (tf TestSearchForm) SetPage(page int) {
	tf.Page = page
}
func (tf TestSearchForm) SetPerPage(perpage int) {
	tf.PerPage = perpage
}
func (tf TestSearchForm) SetSort(sort []string) {
	tf.sort = sort
}
func (tf TestSearchForm) AddSort(field string, asc bool) {
	order := ""
	if asc {
		order = "+"
	} else {
		order = "-"
	}
	tf.sort = append(tf.sort, fmt.Sprintf("%s%s", order, field))
}
func (tf TestSearchForm) SetFilters(filters map[string][]string) {
	tf.filters = filters
}
func (tf TestSearchForm) SetFilter(k string, v []string) {
	tf.filters[k] = v
}
func (tf TestSearchForm) AddFilter(key string, values []string) {
	for _, v := range values {
		tf.filters[key] = append(tf.filters[key], fmt.Sprintf("%v", v))
	}
}
func (tf TestSearchForm) SetKeyword(keyword string) {
	tf.SearchPhrase = keyword
}
func (tf TestSearchForm) KeywordEmbedding() []float32 {
	return tf.embedding
}
func (tf TestSearchForm) RAGUUIDs() []string {
	return tf.ragUUIDs
}
func (tf TestSearchForm) SetRAGUUIDs(ragUUIDs []string) {
	tf.ragUUIDs = ragUUIDs
}

// type Pager interface {
// 	Offset() int
// 	Limit() int
// 	OrderBy() []string

// 	SetPage(page int)
// 	SetPerPage(perpage int)
// 	SetOrders(ol []SortModel)
// 	SetOrder(field string, asc bool)
// }

// type Filterer interface {
// 	Filters() map[string][]string
// 	SetFilters(filters map[string][]string)
// 	SetFilter(k string, v []string)
// 	AddFilter(k string, v string)
// }

// type Searcher interface {
// 	Keyword() string
// 	SetKeyword(keyword string)
// }

func TestList(t *testing.T) {
	conn, err := DefaultSqlxDBConn()
	if err != nil {
		t.Fatal(err.Error())
	}

	repo := New(conn, "test_table2", Config{DebugMode: true}, Table2{})

	sf := TestSearchForm{}
	// sf.filters = make(map[string][]string)
	// sf.SetFilter("name", []string{"a"})

	// base := repo.SelectClause().Columns(repo.TInfo.Columns...).Columns("test_table2.misc").LeftJoin("test_table2 ON test_table2.id = test_table.table2Id")

	result, total, err := repo.List(repo.SelectClause(), sf)

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Printf("Total: %d\n", total)
	for _, v := range result {
		fmt.Printf("%+v\n", v)
	}
}

type ArrayView struct {
	Table2
	Table1List Table1List `db:"testTable" json:"testTable"`
}

func TestArray(t *testing.T) {
	conn, err := DefaultSqlxDBConn()
	if err != nil {
		t.Fatal(err.Error())
	}

	repo := New(conn, "test_table2", Config{DebugMode: true}, Table2{})

	sub := NewSelectClause().From("test_table").Column(ARRAYAGG(JSONObject("test_table", Table1{}))).Add(ExprValue("test_table.table2Id = test_table2.id"))

	result, total, err := repo.List(NewSelectClause().Columns("*").Column(Alias(sub, "testTable")), TestSearchForm{})
	// total, err := repo.List(sub, TestSearchForm{}, &Table1List{})

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Printf("Total: %d\n", total)
	for _, v := range result {
		fmt.Printf("%+v\n", v)
	}

}

func TestDelete(t *testing.T) {
	conn, err := DefaultSqlxDBConn()
	if err != nil {
		t.Fatal(err.Error())
	}
	testcount := 10
	conds := make([]Conditions, testcount)

	// for i := 0; i < testcount; i++ {
	// 	conds[i] = NewConditions().AddEq("id", i+1).AddEq("code", i+1)
	// }

	tableName := "testing"
	repo := New(conn, tableName, Config{}, Base{})

	txRepo, err := repo.TxRepo()
	if err != nil {
		t.Fatal(err.Error())
	}

	if err := txRepo.BatchDelete(conds); err != nil {
		t.Fatal(err)
	}

	if err := txRepo.Commit(); err != nil {
		t.Fatal(err)
	}

}

func TestSoftDelete(t *testing.T) {
	// conn, err := DefaultSqlxDBConn()
	// if err != nil {
	// t.Fatal(err.Error())
	// }
	// repo := New(conn, "test_table2", Config{DebugMode: true}, Table2{})

	// repo.SoftDelete(repo.DeleteClause().AddEq("id", 999))
}

func TestCreate(t *testing.T) {
	conn, err := DefaultSqlxDBConn()
	if err != nil {
		t.Fatal(err.Error())
	}
	repo := New(conn, "test_table2", Config{DebugMode: true}, Table2{})

	repo.Create(map[string]any{"misc": "test"})
}

func TestUpdate(t *testing.T) {
	conn, err := DefaultSqlxDBConn()
	if err != nil {
		t.Fatal(err.Error())
	}
	repo := New(conn, "test_table2", Config{DebugMode: true}, Table2{})

	repo.Update(repo.UpdateClause().AddEq("id", 4), map[string]any{"misc": "5"})
}

type Others struct {
	Detail string `db:"detail" json:"detail"`
}

type Sample struct {
	ID       int            `db:"id" json:"id"`
	Name     string         `db:"name" json:"name"`
	UserInfo map[string]any `db:"user" json:"user" dbfield:"sub"`
	Other    Others         `db:"other" json:"other"`
	Another  *Others        `db:"another" json:"another"`
	Others
	CreatedAt time.Time  `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time  `db:"updatedAt" json:"updatedAt"`
	DeletedAt *time.Time `db:"deletedAt" json:"deletedAt"` //that means we are unable to get the deleted at
	// unix.DeletedAt `db:"deletedAt" json:"deletedAt"` //that means we are unable to get the deleted at
}

func TestPrepare(t *testing.T) {
	//need to see if Sample is a SoftDeleter
	// var same any = Sample{}
	type Asd struct {
		repo *Repo[Sample]
	}

	conn, _ := DefaultSqlxDBConn()
	r := &Asd{
		repo: New(conn, "HiHi", Config{DebugMode: true}, Sample{}),
	}
	// b := r.repo.SelectClause()

	// r.repo.GetAll(b.AddEq("name", "asd"))
	// fmt.Println("done")

	// r.repo.GetAll(r.repo.SelectClause().AddEq("testing", 1))

	// r.repo.SoftDelete(r.repo.DeleteClause().AddEq("id", 122))

	r.repo.Update(r.repo.UpdateClause().AddEq("lol", 23213), map[string]any{"somedata": "do something"})
}
