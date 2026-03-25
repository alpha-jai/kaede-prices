package mysqlrepo

import (
	sq "github.com/Masterminds/squirrel"
)

const DefaultUpdatedAtKey = "updatedAt"

type UpdateBuilder interface {
	Set(column string, value any) UpdateBuilder
	SetMap(clauses map[string]any) UpdateBuilder
	Table(table string) UpdateBuilder
	Limit(limit uint64) UpdateBuilder
	Offset(offset uint64) UpdateBuilder
	OrderBy(orderBys ...string) UpdateBuilder

	WhereBuilder[UpdateBuilder]
	Builder
}

// UpdateClause .
type UpdateClause struct {
	sq.UpdateBuilder
	*WhereClause[UpdateBuilder]
}

func (u UpdateClause) Set(column string, value any) UpdateBuilder {
	u.UpdateBuilder = u.UpdateBuilder.Set(column, value)
	u.WhereClause.Builder = u
	return u
}
func (u UpdateClause) SetMap(clauses map[string]any) UpdateBuilder {
	u.UpdateBuilder = u.UpdateBuilder.SetMap(clauses)
	u.WhereClause.Builder = u
	return u
}
func (u UpdateClause) Table(table string) UpdateBuilder {
	u.UpdateBuilder = u.UpdateBuilder.Table(table)
	u.WhereClause.Builder = u
	return u
}

func (u UpdateClause) Limit(limit uint64) UpdateBuilder {
	u.UpdateBuilder = u.UpdateBuilder.Limit(limit)
	u.WhereClause.Builder = u
	return u
}
func (u UpdateClause) Offset(offset uint64) UpdateBuilder {
	u.UpdateBuilder = u.UpdateBuilder.Offset(offset)
	u.WhereClause.Builder = u
	return u
}

func (u UpdateClause) OrderBy(orderBys ...string) UpdateBuilder {
	u.UpdateBuilder = u.UpdateBuilder.OrderBy(orderBys...)
	u.WhereClause.Builder = u
	return u
}

func (u UpdateClause) ToSql() (string, []any, error) {
	for _, c := range u.Conditions() {
		u.UpdateBuilder = u.UpdateBuilder.Where(c)
	}

	return u.UpdateBuilder.ToSql()
}

func NewUpdateClause() UpdateBuilder {
	uc := UpdateClause{
		UpdateBuilder: sq.UpdateBuilder(sq.StatementBuilder),
		WhereClause:   NewWhereClause[UpdateBuilder](),
	}
	uc.Builder = uc
	return uc
}
