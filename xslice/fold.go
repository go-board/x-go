package xslice

func FoldLeft(a []interface{}, fn func(left, right interface{}) interface{}, initial interface{}) interface{} {
	for _, x := range a {
		initial = fn(initial, x)
	}
	return initial
}

func FoldRight(a []interface{}, fn func(left, right interface{}) interface{}, initial interface{}) interface{} {
	n := len(a)
	for i := n - 1; i >= 0; i-- {
		initial = fn(initial, a[i])
	}
	return initial
}
