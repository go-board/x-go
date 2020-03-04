package xsort

import (
	"sort"

	"github.com/go-board/x-go/types"
)

func SearchUint(a []uint, x uint) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func SearchUint64(a []uint64, x uint64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func SearchUint32(a []uint32, x uint32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func SearchUint16(a []uint16, x uint16) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func SearchUint8(a []uint8, x uint8) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func SearchInt(a []int, x int) int {
	return sort.SearchInts(a, x)
}

func SearchInt64(a []int64, x int64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func SearchInt32(a []int32, x int32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func SearchInt16(a []int16, x int16) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func SearchInt8(a []int8, x int8) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func SearchString(a []string, x string) int {
	return sort.SearchStrings(a, x)
}

func SearchFloat64(a []float64, x float64) int {
	return sort.SearchFloat64s(a, x)
}

func SearchFloat32(a []float32, x float32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func SearchComparable(a []types.Comparable, x types.Comparable) int {
	return sort.Search(len(a), func(i int) bool { return a[i].Compare(x) >= types.OrderingEqual })
}
