package xslice

func Map(a []interface{}, fn func(v interface{}) interface{}) []interface{} {
	b := make([]interface{}, len(a))
	for i, x := range a {
		b[i] = fn(x)
	}
	return b
}
