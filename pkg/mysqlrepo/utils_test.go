package mysqlrepo

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExtractDBTag(t *testing.T) {
	type ShareCreate struct {
		CreatedBy string `db:"created_by"`
		CreatedAt int64  `db:"created_at"`
		Dummy     string
	}
	type ShareUpdate struct {
		UpdatedBy string    `db:"updated_by"`
		UpdatedAt time.Time `db:"updated_at"`
		DummyU    string
	}
	type Hi struct {
		ID       int    `json:"id" db:"id"`
		Name     string `db:"name"`
		Password string `db:"password"`
		Nope     string
		Ptr      *string    `db:"ptr"`
		TestAt   *time.Time `db:"testAt"`
		EscAt    *time.Time `db:"EsctAt"`
		*ShareCreate
		ShareUpdate
		SomeUpdate ShareUpdate `db:"someUpdate"`
	}
	tn := time.Now()
	ptrString := "sfsdf"
	hi := Hi{
		ID:       100,
		Name:     "asdasd",
		Password: "pqwe",
		Nope:     "asdqwe",
		Ptr:      &ptrString,
		TestAt:   &tn,
		ShareCreate: &ShareCreate{
			CreatedBy: "TEST-USER",
			CreatedAt: time.Now().Unix(),
		},
		ShareUpdate: ShareUpdate{
			UpdatedBy: "TEST-USER-UPDATE",
			UpdatedAt: time.Now(),
		},
		SomeUpdate: ShareUpdate{
			UpdatedBy: "someone",
			UpdatedAt: time.Now(),
			DummyU:    "hahahah",
		},
	}

	asd := ExtractDBTag(hi)
	fmt.Println(asd)
}
func TestExtractDBTagSimple(t *testing.T) {
	type GrandChild struct {
		GrandChildName string `db:"grandChildName"`
	}
	type Child struct {
		ChildName   string      `db:"childName"`
		GrandChildA GrandChild  `db:"grandChildA"`
		GrandChildB GrandChild  `db:"grandChildB" dbfield:"sub"`
		GrandChildC *GrandChild `db:"grandChildC"`
		GrandChildD *GrandChild `db:"grandChildC"  dbfield:"sub"`
	}
	type Simple struct {
		Name   string `db:"name"`
		ChildA Child  `db:"childA"`
		ChildB *Child `db:"childB"`
		*Child
		ChildC    Child `db:"childC" dbfield:"sub"`
		ChildD    Child
		CreatedAt time.Time  `db:"createdAt"`
		UpdatedAt *time.Time `db:"updatedAt"`
		DeletedAt *time.Time `db:"deletedAt"`
	}
	tn := time.Now()
	simple := Simple{
		Name: "test",
		ChildA: Child{
			ChildName: "Child A",
		},
		ChildB: &Child{
			ChildName: "Child B",
		},
		Child: &Child{
			ChildName: "Child Embedded",
		},
		ChildC: Child{
			ChildName: "Child C",
		},
		ChildD: Child{
			ChildName: "Child C",
		},
		CreatedAt: tn,
		UpdatedAt: &tn,
	}

	asd := ExtractDBTag(simple)
	fmt.Println(asd)
}

func TestRedisKey(t *testing.T) {
	RedisKeyJWT := "auth-service-jwt:%v:%v:%v"
	expected := "auth-service-jwt:hihi:1:2.2"
	result := RedisKey(RedisKeyJWT, []interface{}{"hihi", 1, 2.2}...)
	fmt.Println(result)
	if result != expected {
		t.Fail()
	}
}

func TestJSON(t *testing.T) {
	type test struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
		Role int    `db:"role"`
	}

	result := JSONObject("user", test{})
	assert.Equal(t, `"id", user.id, "name", user.name, "role", user.role`, result)
}
