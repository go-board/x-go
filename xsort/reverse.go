package xsort

func ReverseInterface(a []interface{}) []interface{} {
	i := 0
	j := len(a)
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	return a
}
