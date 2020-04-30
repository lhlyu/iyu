package util

import "github.com/lhlyu/yutil/v2"

func Md5(v string) string {
	if v == "" {
		return ""
	}
	return yutil.Md5.Encode(v)
}
