package util

import "strings"

func Concat(s ...string) string {
	return strings.Join(s, "")
}
