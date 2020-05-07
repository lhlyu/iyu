package main

import (
	"fmt"
	"github.com/lhlyu/iyu/dao/po"
	"reflect"
)

// 快速生成 赋值
func main() {
	gen(po.Category{}, "v")
}

func gen(v interface{}, prefix string) {
	tp := reflect.TypeOf(v)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	for i := 0; i < tp.NumField(); i++ {
		fmt.Printf("%s: %s.%s,\n", tp.Field(i).Name, prefix, tp.Field(i).Name)
	}
}
