package mysqlrepo

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
)

func TestWhereAndOr(t *testing.T) {

	stmt, args, _ := NewSelectClause().From("table").Columns("test").AddOr(sq.Eq{"id": 11}, sq.Eq{"dob": "today"}).AddOr(sq.Eq{"dob": "tomorrow"}, sq.Eq{"dob": "tomorrow"}).ToSql()
	fmt.Println(stmt)
	fmt.Println(args)
	assert.Equal(t, "SELECT test FROM table WHERE (id = ? OR dob = ?) AND (dob = ? OR dob = ?)", stmt)
	assert.EqualValues(t, []any{11, "today", "tomorrow", "tomorrow"}, args)

}
