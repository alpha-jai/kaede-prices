package mysqlrepo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Role   int    `db:"role"`
	Status int    `db:"status"`
}

type TestItem struct {
	ID        int    `db:"id"`
	Type      string `db:"type"`
	CreatedAt int64  `db:"createdAt"`
	UpdatedAt int64  `db:"updatedAt"`
	DeletedAt int64  `db:"deletedAt"`
	CreatedBy int    `db:"createdBy"`
}

func TestSelectClauseSQLStm(t *testing.T) {
	s := NewSelectClause().From("test_table").Columns("a", "b", "c").Column("count(?) as total", "tester").AddEq("name", "hihi").AddEq("counter", 1)
	stm, val, err := s.ToSql()
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, "SELECT a, b, c, count(?) as total FROM test_table WHERE name = ? AND counter = ?", stm)
	assert.Equal(t, []interface{}{"tester", "hihi", 1}, val)

	sb := NewSelectClause()
	sub := NewSelectClause().From("user").Columns(JSONObject("user", TestUser{})).AddEq("item.createdBy", "user.id")

	alias := Alias(sub, "tester")
	sb = sb.From("item").Columns("item.type", "item.createdAt", "item.updatedAt", "item.deletedAt", "locale.name", "locale.description").
		Column(alias).AddEq("locale.lang", "en").LeftJoin("locale on item.id = locale.itemId").Limit(10).OrderBy("item.ai asc")

	stmt, args, _ := sb.ToSql()

	assert.Equal(t, `SELECT item.type, item.createdAt, item.updatedAt, item.deletedAt, locale.name, locale.description, (SELECT JSON_OBJECT("status", user.status, "id", user.id, "name", user.name, "role", user.role) FROM user WHERE item.createdBy = ?) AS tester FROM item LEFT JOIN locale on item.id = locale.itemId WHERE locale.lang = ? ORDER BY item.ai asc LIMIT 10`, stmt)
	assert.Equal(t, []interface{}{"user.id", "en"}, args)

	items := NewSelectClause().From("item").Columns(ARRAYAGG(JSONObject("item", TestItem{}))).AddEq("item.createdBy", "user.id")
	stmt, args, _ = items.ToSql()
	fmt.Println(stmt, args)
	sb = NewSelectClause().From("user").Column("*").Column(Alias(items, "items")).AddEq("user.id", 1)
	stmt, args, _ = sb.ToSql()
	fmt.Println(stmt, args)
}

func TestSelectClone(t *testing.T) {
	s := NewSelectClause()
	s = s.From("test_table")
	s = s.Columns("a")
	s = s.AddEq("hihi", 1)
	stmt, args, err := s.ToSql()

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(stmt, args)

	newq := s.Clone()
	newq = newq.AddEq("hihi", 1)

	stmt, args, err = newq.ToSql()

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(stmt, args)

	stmt, args, err = s.ToSql()

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(stmt, args)
}
