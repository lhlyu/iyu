package util

import "strconv"

func StringSlinceToIntSlince(ss []string) []int {
	var slince []int
	for _, s := range ss {
		i, e := strconv.Atoi(s)
		if e != nil {
			continue
		}
		slince = append(slince, i)
	}
	return slince
}
