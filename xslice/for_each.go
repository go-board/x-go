package xslice

func ForEach(a []interface{}, fn func(x interface{})) {
	for _, x := range a {
		fn(x)
	}
}

func ForEachString(a []string, fn func(x string)) {
	for _, x := range a {
		fn(x)
	}
}

func ForEachInt(a []int, fn func(x int)) {
	for _, x := range a {
		fn(x)
	}
}

func ForEachInt64(a []int64, fn func(x int64)) {
	for _, x := range a {
		fn(x)
	}
}
