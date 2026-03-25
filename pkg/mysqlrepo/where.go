package mysqlrepo

import (
	sq "github.com/Masterminds/squirrel"
)

type WhereBuilder[T any] interface {
	Add(s sq.Sqlizer) T
	AddAnd(s ...sq.Sqlizer) T
	AddEq(field string, value any) T
	AddGt(field string, value any) T
	AddGtEq(field string, value any) T
	AddIn(field string, value any) T
	AddLike(field string, value any) T
	AddLt(field string, value any) T
	AddLtEq(field string, value any) T
	AddNotEq(field string, value any) T
	AddNotLike(field string, value any) T
	AddNotIn(field string, value any) T
	AddOr(s ...sq.Sqlizer) T

	Conditions() Conditions
	SetConditions(conds Conditions) T
}

func NewWhereClause[T any]() *WhereClause[T] {
	return &WhereClause[T]{
		conds: Conditions{},
	}
}

type WhereClause[T any] struct {
	conds   Conditions
	Builder T
}

type Conditions []sq.Sqlizer

func (w WhereClause[T]) Conditions() Conditions {
	return w.conds
}

func (w *WhereClause[T]) SetConditions(conds Conditions) T {
	w.conds = conds
	return w.Builder
}

func (w *WhereClause[T]) Add(s sq.Sqlizer) T {
	w.conds = append(w.conds, s)
	return w.Builder
}

func (w *WhereClause[T]) AddAnd(s ...sq.Sqlizer) T {
	w.conds = append(w.conds, append(sq.And{}, s...))
	return w.Builder
}

func (w *WhereClause[T]) AddEq(field string, value any) T {
	w.conds = append(w.conds, sq.Eq{field: value})
	return w.Builder
}

func (w *WhereClause[T]) AddGt(field string, value any) T {
	w.conds = append(w.conds, sq.Gt{field: value})
	return w.Builder
}

func (w *WhereClause[T]) AddGtEq(field string, value any) T {
	w.conds = append(w.conds, sq.GtOrEq{field: value})
	return w.Builder
}

func (w *WhereClause[T]) AddIn(field string, value any) T {
	w.conds = append(w.conds, sq.Eq{field: value})
	return w.Builder
}

func (w *WhereClause[T]) AddLike(field string, value any) T {
	w.conds = append(w.conds, sq.Like{field: value})
	return w.Builder
}

func (w *WhereClause[T]) AddLt(field string, value any) T {
	w.conds = append(w.conds, sq.Lt{field: value})
	return w.Builder
}

func (w *WhereClause[T]) AddLtEq(field string, value any) T {
	w.conds = append(w.conds, sq.LtOrEq{field: value})
	return w.Builder
}

func (w *WhereClause[T]) AddNotEq(field string, value any) T {
	w.conds = append(w.conds, sq.NotEq{field: value})
	return w.Builder
}

func (w *WhereClause[T]) AddNotLike(field string, value any) T {
	w.conds = append(w.conds, sq.NotLike{field: value})
	return w.Builder
}

func (w *WhereClause[T]) AddNotIn(field string, value any) T {
	w.conds = append(w.conds, sq.NotEq{field: value})
	return w.Builder
}

func (w *WhereClause[T]) AddOr(s ...sq.Sqlizer) T {
	w.conds = append(w.conds, append(sq.Or{}, s...))
	return w.Builder
}
