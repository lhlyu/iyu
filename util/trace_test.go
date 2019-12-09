package util

import (
	"fmt"
	"testing"
)

type A struct {
}

func (*A) Insert() {
	fmt.Println(RunFuncName(3))
}

func TestGetGID(t *testing.T) {
	a := &A{}
	a.Insert()
}
