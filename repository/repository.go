package repository

import (
	"bytes"
	"reflect"
	"strings"
)

type dao struct {
}

func NewDao() *dao {
	return &dao{}
}

func (*dao) createQuestionMarksForBatch(v ...[]interface{}) {
	//qNum := len(v)
	//maxLength := 0
	//for i := range v {
	//    length := len(v[i])
	//    if length > maxLength{
	//        maxLength = length
	//    }
	//}
	//for i := 0;i < maxLength ;i ++{
	//
	//}
}

// Benchmark-4   	10000000	       161 ns/op  - 10
// Benchmark-4   	 3000000	       448 ns/op  - 100
// Benchmark-4   	 1000000	      1371 ns/op  - 1000
// 可以使用sqlx.In()代替
// Benchmark-4   	 2000000	       782 ns/op  - 10    sqlx.In()
func (*dao) createQuestionMarks(length int) string {
	if length == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	buf.WriteString("?")
	buf.WriteString(strings.Repeat(",?", length-1))
	return buf.String()
}

// any type slice convert to interface slice
// BenchmarkSprintf-4   	 5000000	       381 ns/op - 5
// BenchmarkSprintf-4   	 2000000	       604 ns/op - 10

// BenchmarkSprintf-4   	 5000000	       289 ns/op - 10  not reflect
func (*dao) convertToInterface(slice interface{}) []interface{} {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		return nil
	}
	sliceLen := val.Len()
	if sliceLen == 0 {
		return nil
	}
	params := make([]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		params[i] = val.Index(i).Interface()
	}
	return params
}

// string slice convert to interface slice
func (*dao) strConvertToInterface(slice []string) []interface{} {
	if len(slice) == 0 {
		return nil
	}
	params := make([]interface{}, len(slice))
	for i, v := range slice {
		params[i] = v
	}
	return params
}

// int slice convert to interface slice
func (*dao) intConvertToInterface(slice []int) []interface{} {
	if len(slice) == 0 {
		return nil
	}
	params := make([]interface{}, len(slice))
	for i, v := range slice {
		params[i] = v
	}
	return params
}
