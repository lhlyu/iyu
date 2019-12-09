package util

import (
	"encoding/json"
	"github.com/kataras/iris/core/errors"
	"log"
	"reflect"
)

func ObjToJsonBytes(v interface{}) []byte {
	if v == nil {
		return nil
	}
	bts, _ := json.Marshal(v)
	return bts
}

func ObjToJsonStr(v interface{}) string {
	if v == nil {
		return ""
	}
	bts, _ := json.Marshal(v)
	return string(bts)
}

func JsonStrToObj(s string, v interface{}) error {
	if v == nil {
		return errors.New("v is nil")
	}
	if err := json.Unmarshal([]byte(s), v); err != nil {
		log.Println("JsonStrToObj", s, err, reflect.TypeOf(v))
		return err
	}
	return nil
}
