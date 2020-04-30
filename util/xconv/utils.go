package xconv

import (
	"reflect"
)

func isFuncValid(fType reflect.Type, inTypes, outTypes []interface{}) bool {
	isValid := func(t reflect.Type, required interface{}) bool {
		if required == nil {
			return true
		}
		if typ, ok := required.(reflect.Type); ok {
			if typ.Kind() == reflect.Interface {
				return t.Implements(typ)
			}
			return typ == t
		}
		if kind, ok := required.(reflect.Kind); ok {
			return kind == t.Kind()
		}
		return true
	}

	if fType.Kind() != reflect.Func {
		return false
	}
	if inTypes != nil {
		if fType.NumIn() != len(inTypes) {
			return false
		}
		for i, t := range inTypes {
			if !isValid(fType.In(i), t) {
				return false
			}
		}
	}
	if outTypes != nil {
		if fType.NumOut() != len(outTypes) {
			return false
		}
		for i, t := range outTypes {
			if !isValid(fType.Out(i), t) {
				return false
			}
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
