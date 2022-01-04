package sqlm

import (
	"context"
	"fmt"
	"testing"
)

func TestGenInsertParam(t *testing.T) {
	value := [][]interface{}{
		{1, 100, 200},
		{2, 200, 400},
		{3, 300, 600},
	}
	_, _ = genInsertParam("file", []string{"file_id", "width", "height"}, value)
}

func TestInsert(t *testing.T) {
	db := New(&Config{DSN: "root@(127.0.0.1:3306)/t?charset=utf8mb4&collation=utf8_unicode_ci"})

	r, err := db.Insert(context.Background(), "values", []string{"name", "age"}, [][]interface{}{
		{"小红", 16},
		//{"小丽", 13},
	})
	if err != nil {
        t.Fatal(err)
	}

	affected, err := r.RowsAffected()
	if err != nil {
        t.Fatal(err)
	}
	fmt.Printf("affected: %v\n", affected)

	id, err := r.LastInsertId()
	if err != nil {
        t.Fatal(err)
	}
	fmt.Printf("lastInsertId: %+v\n", id)
}
