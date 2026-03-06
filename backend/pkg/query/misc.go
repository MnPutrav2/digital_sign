package query

import (
	"fmt"
	"reflect"
	"strings"

	_ "github.com/lib/pq"
)

func (q *InitQuery[T]) Where(que string, arg ...string) *InitQuery[T] {

	cut := strings.Split(que, ",")
	if len(cut) == 1 {
		q.query += fmt.Sprintf(" WHERE %s = '%s'", que, arg[0])
	} else {
		q.query += fmt.Sprintf(" WHERE %s = '%s'", cut[0], arg[0])
		for i := range cut {
			if i == 0 {
				continue
			}
			q.query += fmt.Sprintf(" AND %s = '%s'", cut[i], arg[i])
		}
	}

	return q
}

func (q *InitQuery[T]) Like(que string, arg string) *InitQuery[T] {
	q.query += fmt.Sprintf("%s LIKE '%s'", que, "%"+arg+"%")
	return q
}

func (q *InitQuery[T]) And() *InitQuery[T] {
	q.query += " AND "
	return q
}

func (q *InitQuery[T]) Join(table, where string) *InitQuery[T] {

	t := reflect.TypeOf(q.model)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("structure")
		if tag == "" {
			continue
		}

		parts := strings.Split(tag, ";")

		for _, p := range parts[0:] {
			if p == "primary key" {
				f := field.Tag.Get("db")
				q.query += fmt.Sprintf(" INNER JOIN %s ON %s.%s = %s.%s", table, strings.ToLower(t.Name()), strings.ToLower(f), table, where)
			}
		}
	}

	return q
}

func (q *InitQuery[T]) Read() {
	fmt.Println(q.query)
}
