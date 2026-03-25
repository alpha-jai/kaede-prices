package mysqlrepo

type Builder interface {
	ToSql() (string, []any, error)
}
