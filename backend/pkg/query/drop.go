package query

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

func (q *Initdb) Rollback(table ...any) {
	fmt.Println("Running rollback... ==================")
	for _, tbl := range table {
		t := reflect.TypeOf(tbl)

		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		if t.Kind() != reflect.Struct {
			continue
		}

		query := fmt.Sprintf(`DROP TABLE %s`, strings.ToLower(t.Name()))
		if _, err := q.db.Exec(query); err != nil {
			if strings.Contains(err.Error(), "does not exist") ||
				strings.Contains(err.Error(), "not found") {
				fmt.Printf("⚠️  Skipping %s: %v\n", strings.ToLower(t.Name()), err)
				continue
			} else {
				fmt.Println(query)
				log.Fatalf("exec %s: %v", strings.ToLower(t.Name()), err)
			}
		}
	}

	fmt.Println("Rollback completed.")
}
