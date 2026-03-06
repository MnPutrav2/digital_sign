package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

type InitQuery[T any] struct {
	db    *sql.DB
	model T
	query string
}

func Init[T any](db *sql.DB) *InitQuery[T] {
	return &InitQuery[T]{db: db}
}

type Initdb struct {
	db *sql.DB
}

func InitDB(db *sql.DB) *Initdb {
	return &Initdb{db: db}
}

func (q *InitQuery[T]) Exec(ctx context.Context) error {

	if _, err := q.db.ExecContext(ctx, q.query); err != nil {
		fmt.Println(q.query)
		return err
	}

	return nil

}

func (q *InitQuery[T]) Find(ctx context.Context) (T, error) {

	fmt.Println(q.query)
	var result T
	rows, err := q.db.QueryContext(ctx, q.query+" LIMIT 1")
	if err != nil {
		var zero T
		return zero, err
	}
	defer rows.Close()

	if !rows.Next() {
		return result, sql.ErrNoRows
	}

	if err := scan(rows, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (q *InitQuery[T]) FindAll(ctx context.Context) ([]T, error) {

	fmt.Println(q.query)
	rows, err := q.db.QueryContext(ctx, q.query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []T
	for rows.Next() {
		var d T

		err = scan(rows, &d)
		if err != nil {
			return nil, err
		}

		data = append(data, d)
	}

	return data, nil
}

func scan(rows *sql.Rows, dest any) error {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
		return errors.New("dest must be pointer to struct")
	}

	v = v.Elem()
	t := v.Type()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	fieldMap := make(map[string]int)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("db")
		if tag == "" {
			continue
		}

		col := strings.Split(tag, ";")[0]
		fieldMap[col] = i
	}

	scanArgs := make([]any, len(columns))
	uuidFields := make(map[int]int)
	dummy := make([]any, len(columns))

	for i, col := range columns {
		if fieldIndex, ok := fieldMap[col]; ok {

			field := v.Field(fieldIndex)

			if field.Type() == reflect.TypeOf(uuid.UUID{}) {
				var tmp []byte
				scanArgs[i] = &tmp
				uuidFields[i] = fieldIndex
			} else {
				scanArgs[i] = field.Addr().Interface()
			}

		} else {
			scanArgs[i] = &dummy[i]
		}
	}

	if err := rows.Scan(scanArgs...); err != nil {
		return err
	}

	for colIndex, fieldIndex := range uuidFields {
		raw := *(scanArgs[colIndex].(*[]byte))

		id, err := uuid.FromBytes(raw)
		if err != nil {
			id, err = uuid.Parse(string(raw))
			if err != nil {
				return err
			}
			v.Field(fieldIndex).Set(reflect.ValueOf(id))
			continue
		}

		v.Field(fieldIndex).Set(reflect.ValueOf(id))
	}

	return nil
}
