package mysqlrepo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteClauseSQLStm(t *testing.T) {
	dc := NewDeleteClause().From("table").AddIn("in_comp", []string{"hello", "world", "!"}).AddGt("gt_comp", 1).AddLt("lt_comp", 2)
	stm, val, _ := dc.ToSql()
	assert.Equal(t, "DELETE FROM table WHERE in_comp IN (?,?,?) AND gt_comp > ? AND lt_comp < ?", stm)
	assert.Equal(t, []interface{}{"hello", "world", "!", 1, 2}, val)
}
