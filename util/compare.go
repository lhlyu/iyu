package util

// 比较int有差异赋值给第一个
func CompareIntSet(first, second *int) {
	if *second == 0 {
		return
	}
	if *first != *second {
		*first = *second
	}
}

func CompareStrSet(first, second *string) {
	if *second == "" {
		return
	}
	if *first != *second {
		*first = *second
	}
}
