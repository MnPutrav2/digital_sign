package query

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (q *InitQuery[T]) Insert(D any) *InitQuery[T] {

	val := reflect.TypeOf(q.model)
	field := val.NumField()
	val2 := reflect.ValueOf(D)

	t := reflect.TypeOf(q.model)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	var args string
	var v string

	for i := range field {
		if i == val.NumField()-1 {
			args += val.Field(i).Tag.Get("db")
			x := val2.Field(i).Interface()
			switch x.(type) {
			case int, bool:
				v += fmt.Sprintf("%d", val2.Field(i).Int())
			case uuid.UUID:
				v += fmt.Sprintf("'%s', ", val2.Field(i).Interface().(uuid.UUID).String())
			case string:
				v += fmt.Sprintf("'%s'", val2.Field(i).String())
			case time.Time:
				v += fmt.Sprintf("'%s'", val2.Field(i).Interface().(time.Time).Format("2006-01-02 15:04:05"))
			}
		} else {
			args += val.Field(i).Tag.Get("db") + ", "
			x := val2.Field(i).Interface()
			switch x.(type) {
			case int, bool:
				v += fmt.Sprintf("%d, ", val2.Field(i).Int())
			case uuid.UUID:
				v += fmt.Sprintf("'%s', ", val2.Field(i).Interface().(uuid.UUID).String())
			case string:
				v += fmt.Sprintf("'%s', ", val2.Field(i).String())
			case time.Time:
				v += fmt.Sprintf("'%s'", val2.Field(i).Interface().(time.Time).Format("2006-01-02 15:04:05"))
			}
		}
	}

	q.query = fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", strings.ToLower(t.Name()), args, v)
	return q
}
