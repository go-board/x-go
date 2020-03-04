package xslice

import (
	"github.com/go-board/x-go/types"
)

func Contains(a []interface{}, x interface{}) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

func ContainsComparable(a []types.Comparable, x types.Comparable) bool {
	for _, v := range a {
		if v.Compare(x) == types.OrderingEqual {
			return true
		}
	}
	return false
}

func ContainsString(a []string, x string) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

func ContainsInt64(a []int64, x int64) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}
