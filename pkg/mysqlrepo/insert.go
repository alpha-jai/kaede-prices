package mysqlrepo

import (
	"strings"

	sq "github.com/Masterminds/squirrel"
)

type InsertBuilder interface {
	Columns(columns ...string) InsertBuilder
	Into(from string) InsertBuilder
	Prefix(sql string, args ...any) InsertBuilder
	PrefixExpr(expr sq.Sqlizer) InsertBuilder
	Select(sb sq.SelectBuilder) InsertBuilder
	SetMap(clauses map[string]any) InsertBuilder
	Suffix(sql string, args ...any) InsertBuilder
	SuffixExpr(expr sq.Sqlizer) InsertBuilder
	Values(values ...any) InsertBuilder
	OnDuplicateKeyUpdateSuffix(fields ...string) InsertBuilder
	OnDuplicateKeyIgnoreSuffix(dummyField string) InsertBuilder

	Builder
}

// InsertClause .
type InsertClause struct {
	sq.InsertBuilder
}

func (ic InsertClause) Columns(columns ...string) InsertBuilder {
	ic.InsertBuilder = ic.InsertBuilder.Columns(columns...)
	return ic
}
func (ic InsertClause) Into(from string) InsertBuilder {
	ic.InsertBuilder = ic.InsertBuilder.Into(from)
	return ic
}

func (ic InsertClause) Prefix(sql string, args ...any) InsertBuilder {
	ic.InsertBuilder = ic.InsertBuilder.Prefix(sql, args...)
	return ic
}
func (ic InsertClause) PrefixExpr(expr sq.Sqlizer) InsertBuilder {
	ic.InsertBuilder = ic.InsertBuilder.PrefixExpr(expr)
	return ic
}
func (ic InsertClause) Select(sb sq.SelectBuilder) InsertBuilder {
	ic.InsertBuilder = ic.InsertBuilder.Select(sb)
	return ic
}
func (ic InsertClause) SetMap(clauses map[string]any) InsertBuilder {
	ic.InsertBuilder = ic.InsertBuilder.SetMap(clauses)
	return ic
}
func (ic InsertClause) Suffix(sql string, args ...any) InsertBuilder {
	ic.InsertBuilder = ic.InsertBuilder.Suffix(sql, args...)
	return ic
}
func (ic InsertClause) SuffixExpr(expr sq.Sqlizer) InsertBuilder {
	ic.InsertBuilder = ic.InsertBuilder.SuffixExpr(expr)
	return ic
}
func (ic InsertClause) Values(values ...any) InsertBuilder {
	ic.InsertBuilder = ic.InsertBuilder.Values(values...)
	return ic
}

// OnDuplicateKeyIgnoreSuffix add a suffix expr to the insert clause which
// update on duplicated fields.
// The parameter fields in here will be fields you want to update when insert
// on duplicated key.
func (ic InsertClause) OnDuplicateKeyUpdateSuffix(fields ...string) InsertBuilder {
	if len(fields) == 0 {
		return ic
	}
	var sb strings.Builder
	sb.WriteString(`ON DUPLICATE KEY UPDATE `)
	for i := range fields {
		sb.WriteString(fields[i])
		sb.WriteString(`=VALUES(`)
		sb.WriteString(fields[i])
		sb.WriteString(`)`)
		if i < len(fields)-1 {
			sb.WriteString(`,`)
		}
	}
	ic.InsertBuilder = ic.InsertBuilder.SuffixExpr(sq.Expr(sb.String()))
	return ic
}

// OnDuplicateKeyIgnoreSuffix add a suffix expr to the insert clause
// which will skip the insert if duplicated fields are found.
// dummyField is any field of the table, it is best to use a field
// which is not unique key or primary key for dummyField.
func (ic InsertClause) OnDuplicateKeyIgnoreSuffix(dummyField string) InsertBuilder {
	if dummyField == "" {
		return ic
	}
	var sb strings.Builder
	sb.WriteString(`ON DUPLICATE KEY UPDATE `)
	sb.WriteString(dummyField)
	sb.WriteString(`=`)
	sb.WriteString(dummyField)

	ic.InsertBuilder = ic.InsertBuilder.SuffixExpr(sq.Expr(sb.String()))
	return ic
}

func NewInsertClause() InsertBuilder {
	return InsertClause{
		InsertBuilder: sq.InsertBuilder(sq.StatementBuilder),
	}
}
