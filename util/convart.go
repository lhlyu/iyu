package util

import "time"

//
func TimeHandler(s time.Time) int64 {
	return s.Unix() * 10e3
}
