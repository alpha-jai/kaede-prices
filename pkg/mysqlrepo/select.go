package mysqlrepo

import (
	sq "github.com/Masterminds/squirrel"
)

type SelectBuilder interface {
	Columns(columns ...string) SelectBuilder
	Column(column any, args ...any) SelectBuilder
	From(from string) SelectBuilder
	CrossJoin(join string, rest ...any) SelectBuilder
	Distinct() SelectBuilder
	FromSelect(from sq.SelectBuilder, alias string) SelectBuilder
	GroupBy(groupBys ...string) SelectBuilder
	Having(pred any, rest ...any) SelectBuilder
	InnerJoin(join string, rest ...any) SelectBuilder
	Join(join string, rest ...any) SelectBuilder
	JoinClause(pred any, args ...any) SelectBuilder
	LeftJoin(join string, rest ...any) SelectBuilder
	Limit(limit uint64) SelectBuilder
	Offset(offset uint64) SelectBuilder
	OrderBy(orderBys ...string) SelectBuilder
	OrderByClause(pred any, args ...any) SelectBuilder
	Prefix(sql string, args ...any) SelectBuilder
	PrefixExpr(expr sq.Sqlizer) SelectBuilder
	RemoveColumns() SelectBuilder
	RemoveLimit() SelectBuilder
	RemoveOffset() SelectBuilder
	RightJoin(join string, rest ...any) SelectBuilder
	Suffix(sql string, args ...any) SelectBuilder
	SuffixExpr(expr sq.Sqlizer) SelectBuilder

	Clone() SelectBuilder

	WhereBuilder[SelectBuilder]
	Builder
}

type SelectClause struct {
	sq.SelectBuilder
	*WhereClause[SelectBuilder]
}

func (s SelectClause) Columns(columns ...string) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.Columns(columns...)
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) From(from string) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.From(from)
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) Column(column any, args ...any) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.Column(column, args...)
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) CrossJoin(join string, rest ...any) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.CrossJoin(join, rest...)
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) Distinct() SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.Distinct()
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) FromSelect(from sq.SelectBuilder, alias string) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.FromSelect(from, alias)
	s.WhereClause.Builder = s
	return s
}
func (s SelectClause) GroupBy(groupBys ...string) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.GroupBy(groupBys...)
	s.WhereClause.Builder = s
	return s
}
func (s SelectClause) Having(pred any, rest ...any) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.Having(pred, rest...)
	s.WhereClause.Builder = s
	return s
}
func (s SelectClause) InnerJoin(join string, rest ...any) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.InnerJoin(join, rest...)
	s.WhereClause.Builder = s
	return s
}
func (s SelectClause) Join(join string, rest ...any) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.Join(join, rest...)
	s.WhereClause.Builder = s
	return s
}
func (s SelectClause) JoinClause(pred any, args ...any) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.JoinClause(pred, args...)
	s.WhereClause.Builder = s
	return s
}
func (s SelectClause) LeftJoin(join string, rest ...any) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.LeftJoin(join, rest...)
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) Limit(limit uint64) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.Limit(limit)
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) Offset(offset uint64) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.Offset(offset)
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) OrderBy(orderBys ...string) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.OrderBy(orderBys...)
	s.WhereClause.Builder = s
	return s
}
func (s SelectClause) OrderByClause(pred any, args ...any) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.OrderByClause(pred, args...)
	s.WhereClause.Builder = s
	return s
}
func (s SelectClause) Prefix(sql string, args ...any) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.Prefix(sql, args...)
	s.WhereClause.Builder = s
	return s
}
func (s SelectClause) PrefixExpr(expr sq.Sqlizer) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.PrefixExpr(expr)
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) RemoveColumns() SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.RemoveColumns()
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) RemoveLimit() SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.RemoveLimit()
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) RemoveOffset() SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.RemoveOffset()
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) RightJoin(join string, rest ...any) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.RightJoin(join, rest...)
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) Suffix(sql string, args ...any) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.Suffix(sql, args...)
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) SuffixExpr(expr sq.Sqlizer) SelectBuilder {
	s.SelectBuilder = s.SelectBuilder.SuffixExpr(expr)
	s.WhereClause.Builder = s
	return s
}

func (s SelectClause) ToSql() (string, []any, error) {
	//we do that here
	for _, c := range s.Conditions() {
		s.SelectBuilder = s.SelectBuilder.Where(c)
	}
	return s.SelectBuilder.ToSql()
}

func (s SelectClause) Clone() SelectBuilder {
	newsc := s
	newsc.WhereClause = NewWhereClause[SelectBuilder]()
	newsc.WhereClause.Builder = newsc
	newsc.WhereClause.SetConditions(s.Conditions())
	return newsc
}

func NewSelectClause() SelectBuilder {
	sc := SelectClause{
		SelectBuilder: sq.SelectBuilder(sq.StatementBuilder),
		WhereClause:   NewWhereClause[SelectBuilder](),
	}
	return sc
}
