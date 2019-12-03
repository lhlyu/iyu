package repository

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type dao struct {
}

func NewDao() *dao {
	return &dao{}
}

func (d *dao) createQuestionMarksForBatch(v ...[]interface{}) (string, []interface{}) {
	if len(v) == 0 {
		return "", nil
	}
	qNum := len(v[0])
	maxLength := len(v)
	qm := d.createQuestionMarks(qNum)
	var params []interface{}
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf(" values(%s)", qm))
	params = append(params, v[0]...)
	for i := 1; i < maxLength; i++ {
		buf.WriteString(",(")
		buf.WriteString(qm)
		buf.WriteString(")")
		params = append(params, v[i]...)
	}
	return buf.String(), params
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
