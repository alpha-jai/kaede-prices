package mysqlrepo

import (
	sq "github.com/Masterminds/squirrel"
)

const DefaultSoftDeleteKey = "deleted_at"

type DeleteBuilder interface {
	From(from string) DeleteBuilder
	Limit(limit uint64) DeleteBuilder
	Offset(offset uint64) DeleteBuilder
	OrderBy(orderBys ...string) DeleteBuilder

	WhereBuilder[DeleteBuilder]
	Builder
}

// DeleteClause .
type DeleteClause struct {
	sq.DeleteBuilder
	*WhereClause[DeleteBuilder]
}

func (dc DeleteClause) From(from string) DeleteBuilder {
	dc.DeleteBuilder = dc.DeleteBuilder.From(from)
	dc.WhereClause.Builder = dc
	return dc
}

func (dc DeleteClause) Limit(limit uint64) DeleteBuilder {
	dc.DeleteBuilder = dc.DeleteBuilder.Limit(limit)
	dc.WhereClause.Builder = dc
	return dc
}

func (dc DeleteClause) Offset(offset uint64) DeleteBuilder {
	dc.DeleteBuilder = dc.DeleteBuilder.Offset(offset)
	dc.WhereClause.Builder = dc
	return dc
}

func (dc DeleteClause) OrderBy(orderBys ...string) DeleteBuilder {
	dc.DeleteBuilder = dc.DeleteBuilder.OrderBy(orderBys...)
	dc.WhereClause.Builder = dc
	return dc
}

func (dc DeleteClause) ToSql() (string, []any, error) {
	for _, c := range dc.Conditions() {
		dc.DeleteBuilder = dc.DeleteBuilder.Where(c)
	}
	return dc.DeleteBuilder.ToSql()
}

func NewDeleteClause() DeleteBuilder {
	dc := DeleteClause{
		DeleteBuilder: sq.DeleteBuilder(sq.StatementBuilder),
		WhereClause:   NewWhereClause[DeleteBuilder](),
	}
	dc.Builder = dc
	return dc
}
