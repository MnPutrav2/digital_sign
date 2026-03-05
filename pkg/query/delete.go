package query

import (
	"fmt"
	"reflect"
	"strings"
)

func (q *InitQuery[T]) Delete(que string, arg ...string) *InitQuery[T] {
	t := reflect.TypeOf(q.model)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	q.query = fmt.Sprintf("DELETE FROM %s", strings.ToLower(t.Name()))
	cut := strings.Split(que, ",")
	if len(cut) == 1 {
		q.query += fmt.Sprintf(" WHERE %s = '%s'", que, arg[0])
	} else {
		q.query += fmt.Sprintf(" WHERE %s = '%s'", strings.TrimSpace(cut[0]), arg[0])
		for i := range cut {
			if i == 0 {
				continue
			}

			q.query += fmt.Sprintf(" AND %s = '%s'", strings.TrimSpace(cut[i]), arg[i])
		}
	}

	return q
}
