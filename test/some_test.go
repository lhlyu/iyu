package test

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

type A struct {
	List []string `json:"list"`
}

func TestSome(t *testing.T) {
	a := &A{
		List: nil,
	}
	bytes, _ := json.Marshal(a)
	fmt.Println(string(bytes))
	sql, params, _ := sqlx.In("select * from demo where b = ? id in (?) and a = ?", 6, []int{1, 2, 3}, 4)
	fmt.Println(sql, params)
}
