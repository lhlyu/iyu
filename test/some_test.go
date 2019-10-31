package test

import (
	"encoding/json"
	"fmt"
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
}
