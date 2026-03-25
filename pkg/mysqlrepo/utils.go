package mysqlrepo

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func extractTag(tag string, structObj interface{}, m map[string]interface{}) {
	fields := reflect.TypeOf(structObj)
	values := reflect.ValueOf(structObj)

	if fields.Kind() == reflect.Ptr {
		fields = fields.Elem()
	}

	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}

	fieldsN := fields.NumField()
	for i := 0; i < fieldsN; i++ {
		field := fields.Field(i)
		fieldName := field.Tag.Get(tag)
		//skip empty db tag
		value := values.Field(i)
		switch value.Kind() {
		case reflect.Struct:
			// check if field is time.Time
			if field.Type == reflect.TypeOf(time.Time{}) {
				m[fieldName] = value.Interface()
				continue
			}
			if fields.Field(i).Anonymous {
				extractTag(tag, value.Interface(), m)
				continue
			}
		case reflect.Ptr:
			if value.IsNil() {
				continue
			}
			if value.Elem().Kind() == reflect.Struct {
				if field.Type.Elem() == reflect.TypeOf(time.Time{}) {
					m[fieldName] = value.Interface()
				} else {
					extractTag(tag, value.Interface(), m)
				}
				continue
			}
			//skip non-struct field without tag field
			if fieldName == "" {
				continue
			}
			m[fieldName] = value.Elem().Interface()
		default:
			if fieldName == "" {
				continue
			}
			m[fieldName] = value.Interface()
		}
	}
}

func extractDBTag(structObj interface{}, m map[string]interface{}) {
	fields := reflect.TypeOf(structObj)
	values := reflect.ValueOf(structObj)

	if fields.Kind() == reflect.Ptr {
		fields = fields.Elem()
	}

	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}

	fieldsN := fields.NumField()
	for i := 0; i < fieldsN; i++ {
		field := fields.Field(i)
		fieldName := field.Tag.Get("db")
		dbField := field.Tag.Get("dbfield")
		//skip empty db tag
		value := values.Field(i)
		switch value.Kind() {
		case reflect.Struct:
			if fields.Field(i).Anonymous {
				extractDBTag(value.Interface(), m)
			} else {
				if fieldName == "" {
					continue
				}
				dbFieldMap := ToDBFieldMap(dbField)
				if !IsDBField(dbFieldMap, "sub") {
					m[fieldName] = value.Interface()
				}
			}
		case reflect.Ptr:
			if value.IsNil() {
				continue
			}
			if value.Elem().Kind() == reflect.Struct {
				if field.Anonymous {
					extractDBTag(value.Interface(), m)
				} else {
					if fieldName == "" {
						continue
					}
					dbFieldMap := ToDBFieldMap(dbField)
					if !IsDBField(dbFieldMap, "sub") {
						m[fieldName] = value.Interface()
					}
				}
				continue
			}
			//skip non-struct field without tag field
			if fieldName == "" {
				continue
			}
			m[fieldName] = value.Elem().Interface()
		default:
			if fieldName == "" {
				continue
			}
			m[fieldName] = value.Interface()
		}
	}
}

// ExtractTag extracts data from a struct with input tag to a map
func ExtractTag(tag string, structObj interface{}) map[string]interface{} {
	m := map[string]interface{}{}

	extractTag(tag, structObj, m)

	return m
}

// ExtractDBTag extracts struct fields with tag db to a map
func ExtractDBTag(structObj interface{}) map[string]interface{} {
	m := map[string]interface{}{}

	extractDBTag(structObj, m)

	return m
}

func RedisKey(source string, values ...interface{}) string {
	return fmt.Sprintf(source, values...)
}

func JSONObject(prefix string, value any) string {
	tags := ExtractDBTag(value)
	result := ""
	for k := range tags {
		if result != "" {
			result += ", "
		}
		result += fmt.Sprintf(`"%s", %s.%s`, k, prefix, k)
	}
	return fmt.Sprintf("JSON_OBJECT(%s)", result)
}

func ARRAYAGG(field string) string {
	return fmt.Sprintf("JSON_ARRAYAGG(%s)", field)
}

func Alias(expr sq.Sqlizer, alias string) sq.Sqlizer {
	return sq.Alias(expr, alias)
}

func StmtConcatWS(seperator string, args ...string) string {
	stmt, _, _ := ExprValue(fmt.Sprintf(`CONCAT_WS('%s', %s)`, seperator, strings.Join(args, ","))).ToSql()
	return stmt
}

func ExprValue(sql string, args ...interface{}) sq.Sqlizer {
	return sq.Expr(sql, args...)
}

func marshalOrders(sorts []string) []string {
	result := []string{}
	for _, v := range sorts {
		field := ""
		order := "asc"
		switch v[0] {
		case '-':
			field = v[1:]
			order = "desc"
		case '+':
			field = v[1:]
		default:
			field = v
		}
		result = append(result, fmt.Sprintf("%s %s", field, order))
	}
	return result
}

// MatchScore creates a MATCH expression for full-text search with an alias
func MatchScore(columns []string, searchTerm string) sq.Sqlizer {
	matchExpr := fmt.Sprintf("MATCH(%s) AGAINST(? IN NATURAL LANGUAGE MODE)", strings.Join(columns, ","))
	return ExprValue(matchExpr, searchTerm)
}

// MatchWhere creates a MATCH expression for use in WHERE clauses
func MatchWhere(columns []string, searchTerm string) sq.Sqlizer {
	matchExpr := fmt.Sprintf("MATCH(%s) AGAINST(? IN NATURAL LANGUAGE MODE)", strings.Join(columns, ","))
	return ExprValue(matchExpr, searchTerm)
}
