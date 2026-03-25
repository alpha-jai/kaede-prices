package mysqlrepo

import (
	"reflect"
	"time"
)

type TimeType int

const (
	None TimeType = iota
	Unix
	Time
	TimePtr
)

func ToTimeType(t reflect.Type) TimeType {
	switch {
	case t == reflect.TypeOf(time.Time{}):
		return Time
	case t == reflect.TypeOf(&time.Time{}):
		return TimePtr
	case t.Kind() == reflect.Int:
		return Unix
	case t.Kind() == reflect.Int64:
		return Unix
	default:
		return None
	}
}

func TimeNow(t TimeType) any {
	switch t {
	case Unix:
		return time.Now().Unix()
	case Time:
		return time.Now().UTC()
	case TimePtr:
		now := time.Now().UTC()
		return &now
	default:
		return nil
	}
}

func TimeEmpty(t TimeType) any {
	switch t {
	case Unix:
		return 0
	case Time:
		return time.Time{}
	case TimePtr:
		return nil
	default:
		return nil
	}
}

func AutoTime(t TimeType) bool {
	switch t {
	case Unix, Time:
		return true
	default:
		return false
	}
}
