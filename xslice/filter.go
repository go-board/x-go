package xslice

func Filter(a []interface{}, fn func(v interface{}) bool) []interface{} {
	b := make([]interface{}, 0)
	for _, x := range a {
		if fn(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterString(a []string, fn func(v string) bool) []string {
	b := make([]string, 0)
	for _, x := range a {
		if fn(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterInt(a []int, fn func(v int) bool) []int {
	b := make([]int, 0)
	for _, x := range a {
		if fn(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterInt64(a []int64, fn func(v int64) bool) []int64 {
	b := make([]int64, 0)
	for _, x := range a {
		if fn(x) {
			b = append(b, x)
		}
	}
	return b
}
