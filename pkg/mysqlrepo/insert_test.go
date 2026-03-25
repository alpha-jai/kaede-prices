package mysqlrepo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertClauseSQLStm(t *testing.T) {
	ic := NewInsertClause().Into("table").SetMap(map[string]interface{}{
		"field": 1,
	})

	stm, val, _ := ic.ToSql()
	assert.Equal(t, "INSERT INTO table (field) VALUES (?)", stm)
	assert.Equal(t, []interface{}{1}, val)
}

func TestInsertClauseOnDuplicateKeyUpdateSuffix(t *testing.T) {
	ic := NewInsertClause().SetMap(map[string]interface{}{
		"field": 1,
	}).OnDuplicateKeyUpdateSuffix("field").Into("table")

	stm, val, _ := ic.ToSql()
	assert.Equal(t, "INSERT INTO table (field) VALUES (?) ON DUPLICATE KEY UPDATE field=VALUES(field)", stm)
	assert.Equal(t, []interface{}{1}, val)
}

func TestBatchInsert(t *testing.T) {
	ic := NewInsertClause().Into("table").SetMap(map[string]interface{}{
		"field": 1,
	}).SetMap(map[string]interface{}{
		"field": 2,
	}).SetMap(map[string]interface{}{
		"field": 3,
	})

	stm, val, _ := ic.ToSql()
	// assert.Equal(t, "INSERT INTO table (field) VALUES (?)", stm)
	// assert.Equal(t, []interface{}{1}, val)
	fmt.Println(stm, val)
}
