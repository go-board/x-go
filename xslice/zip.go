package xslice

import (
	"fmt"
)

type ZipItem struct {
	Left  interface{}
	Right interface{}
}

func Zip(a []interface{}, b []interface{}) []*ZipItem {
	if len(a) != len(b) {
		panic(fmt.Sprintf("len a is %d, len b is %d", len(a), len(b)))
	}
	items := make([]*ZipItem, 0, len(a))
	for i, x := range a {
		items = append(items, &ZipItem{
			Left:  x,
			Right: b[i],
		})
	}
	return items
}
