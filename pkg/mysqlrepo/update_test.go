package mysqlrepo

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
)

func TestUpdateClauseSQLStm(t *testing.T) {
	uc := NewUpdateClause().Table("table").SetMap(map[string]interface{}{
		"field1": 1,
		"field2": "2",
	}).AddIn("in_comp", []string{"hello", "world", "!"}).AddGt("gt_comp", 1).AddLt("lt_comp", 2)
	stm, val, _ := uc.ToSql()
	assert.Equal(t, "UPDATE table SET field1 = ?, field2 = ? WHERE in_comp IN (?,?,?) AND gt_comp > ? AND lt_comp < ?", stm)
	assert.Equal(t, []interface{}{1, "2", "hello", "world", "!", 1, 2}, val)
}

func BenchmarkUpdateClause(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uc := NewUpdateClause().Table("table").
			// SetMap(map[string]interface{}{"field1": 1, "field2": "2"}).
			Set("field", 1).Set("field2", 2).
			AddIn("in_comp", []string{"hello", "world", "!"}).AddGt("gt_comp", 1).AddLt("lt_comp", 2)
		uc.ToSql()
	}
}

func BenchmarkOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sq.Update("table").Set("field", 1).Set("field2", 2).Where("in_comp", []string{"hello", "world", "!"}).Where(sq.Gt{"gt_comp": 1}).Where(sq.Lt{"lt_comp": 2}).ToSql()
	}
}
