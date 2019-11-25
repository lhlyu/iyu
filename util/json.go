package util

import (
	"encoding/json"
	"log"
)

func ObjToJsonStr(v interface{}) string {
	if v == nil {
		return ""
	}
	bts, _ := json.Marshal(v)
	return string(bts)
}

func JsonStrToObj(s string, v interface{}) error {
	if v == nil {
		return nil
	}
	if err := json.Unmarshal([]byte(s), v); err != nil {
		log.Println("JsonStrToObj", s, err)
		return err
	}
	return nil
}
