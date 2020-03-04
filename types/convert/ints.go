package convert

import (
	"strconv"

	"github.com/go-board/x-go/types"
)

func IntSliceToInt64Slice(a types.IntSlice) types.Int64Slice {
	b := make(types.Int64Slice, len(a))
	for i, x := range a {
		b[i] = int64(x)
	}
	return b
}

func Int64SliceToIntSlice(a types.Int64Slice) types.IntSlice {
	b := make(types.IntSlice, len(a))
	for i, x := range a {
		b[i] = int(x)
	}
	return b
}

func IntSliceToAnySlice(a types.IntSlice) types.AnySlice {
	any := make(types.AnySlice, len(a))
	for i, x := range a {
		any[i] = x
	}
	return any
}

func Int64SliceToAnySlice(a types.Int64Slice) types.AnySlice {
	b := make(types.AnySlice, len(a))
	for i, x := range a {
		b[i] = x
	}
	return b
}

func IntSliceToStringSlice(a types.IntSlice) types.StringSlice {
	b := make(types.StringSlice, len(a))
	for i, x := range a {
		b[i] = strconv.Itoa(x)
	}
	return b
}

func Int64SliceToStringSlice(a types.Int64Slice) types.StringSlice {
	b := make(types.StringSlice, len(a))
	for i, x := range a {
		b[i] = strconv.FormatInt(x, 10)
	}
	return b
}

func StringSliceToIntSlice(a types.StringSlice) (b types.IntSlice, hasErr bool) {
	for _, x := range a {
		i, err := strconv.Atoi(x)
		if err != nil {
			hasErr = true
			continue
		}
		b = append(b, i)
	}
	return
}

func StringSliceToInt64Slice(a types.StringSlice) (b types.Int64Slice, hasErr bool) {
	for _, x := range a {
		i, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			hasErr = true
			continue
		}
		b = append(b, i)
	}
	return
}
