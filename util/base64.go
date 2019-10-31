package util

import (
	"encoding/base64"
	"encoding/json"
)

func Base64Encode(bytes []byte) string {
	result := base64.StdEncoding.EncodeToString(bytes)
	return result
}

func Base64EncodeString(data string) string {
	return Base64Encode([]byte(data))
}

func Base64EncodeObj(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return Base64Encode(bytes)
}

func Base64Decode(data string) []byte {
	bytes, _ := base64.StdEncoding.DecodeString(data)
	return bytes
}

func Base64DecodeString(data string) string {
	bytes := Base64Decode(data)
	return string(bytes)
}

func Base64DecodeObj(data string, v interface{}) {
	bytes := Base64Decode(data)
	json.Unmarshal(bytes, v)
}
