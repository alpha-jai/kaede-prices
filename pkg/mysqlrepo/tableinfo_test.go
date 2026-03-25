package mysqlrepo

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type Common struct {
	Bye string `db:"bye"`
}

type Base struct {
	Hello     string  `db:"hello" dbfield:"pk"`
	Something *string `db:"something" dbfield:"search,unique"`
	Common
}

type TestEntity struct {
	CreatedAt int64 `db:"createdAt" json:"createdAt"`
	UpdatedAt int64 `db:"updatedAt" json:"updatedAt"`
	DeletedAt int64 `db:"deletedAt" json:"deletedAt"`
}

func (e TestEntity) ConvertValue(t time.Time) int64 {
	return t.Unix()
}

func (e TestEntity) GetCreatedAtTag() string {
	return "createdAt"
}

func (e TestEntity) GetUpdatedAtTag() string {
	return "updatedAt"
}

func (e TestEntity) GetDeletedAtTag() string {
	return "deletedAt"
}

func (e TestEntity) GetCreatedAtType() reflect.Kind {
	return reflect.Int64
}

func (e TestEntity) GetUpdatedAtType() reflect.Kind {
	return reflect.Int64
}

func (e TestEntity) GetDeletedAtType() reflect.Kind {
	return reflect.Int64
}

type TestDefaultEntity struct {
	// EntityUnix
}

type Child struct {
	Hihi string `db:"hihi"`
	Base
}

type NestedChild struct {
	Hihi       string `db:"hihi"`
	BaseStruct Base
}

func TestPrepareInfo(t *testing.T) {
	a := PrepareDBTableInfo("testing", TestDefaultEntity{})
	b := PrepareDBTableInfo("testing", TestEntity{})
	fmt.Println(a.Columns)
	fmt.Println(b.Columns)
}

func TestPrepareInfoPtr(t *testing.T) {
	a := PrepareDBTableInfo("testing", &TestDefaultEntity{})
	b := PrepareDBTableInfo("testing", &TestEntity{})
	fmt.Println(a)
	fmt.Println(b)
}

func TestPrepareInfoUpdater(t *testing.T) {
	type Update struct {
		Name      string     `db:"name" json:"name"`
		CreatedAt time.Time  `db:"createdAt" json:"createdAt"`
		UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt"`
		DeletedAt int64      `db:"deleted_at" json:"deletedAt"`
	}

	a := PrepareDBTableInfo("testing", &Update{})
	fmt.Println(a.Columns)
}

func TestFilterKeys(t *testing.T) {
	source := map[string]interface{}{
		"id":         13,
		"name":       "Testing",
		"status":     1,
		"created_at": time.Now(),
		"updated_at": "",
	}
	reserveKeys := []string{"name", "status", "updated_at"}
	result := FilterKeys(source, reserveKeys)
	if len(result) != len(reserveKeys) {
		t.Error("result length not match")
		return
	}
	for k, v := range result {
		oriVal, ok := source[k]
		if !ok {
			t.Error("reserved key has been accidentally filtered out")
			return
		}
		if oriVal != v {
			t.Error("reserved key value has been accidentally changed")
			return
		}
	}
}

func TestSubQuery(t *testing.T) {
	type SubUserModel struct {
		ID     int64  `db:"id" dbfield:"ai"`
		Name   string `db:"name" dbfield:"search"`
		Status int    `db:"status"`
	}

	type Model struct {
		ID    int64        `db:"id" dbfield:"ai"`
		User  SubUserModel `db:"sub" dbfield:"search,alias=xxx,sub=???"` //sub=user meaning a subquery to user table
		Count int          `db:"total" dbfield:"alias=count(*)"`         //maybe read json as well? if just json then its ok. but we are not it is a subquery
	}

	jsonObj := JSONObject("user", SubUserModel{})
	sub := sq.Select(jsonObj).From("user").Where("table1.id = user.id")

	result := PrepareDBTableInfo("testing", &Model{})
	stmt, _, _ := sq.Alias(sub, "model").ToSql()
	result.Columns = append(result.Columns, stmt)
	fmt.Println(result.Columns)
}

func TestColumnMap(t *testing.T) {

	type Model struct {
		ID       int64  `db:"id" dbfield:"ai"`
		Count    int    `db:"total" dbfield:"alias=count(*)" table:""` //maybe read json as well? if just json then its ok. but we are not it is a subquery
		JoinName string `db:"joinName" dbfield:"name" table:"join_table"`
	}

	result := PrepareDBTableInfo("testing", &Model{})
	fmt.Println(result.ColumnMap)
}

func TestGetEntityType(t *testing.T) {
	// u := unix.Entity{}

	tn := time.Now()
	var u any = tn

	switch ast := u.(type) {
	case string:
		fmt.Println("String")
	case int64:
		fmt.Println("int64")
	case *time.Time:
		fmt.Println("time.time")
	default:
		fmt.Println("default", ast)
	}
}
