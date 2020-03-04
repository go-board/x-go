package xslice

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-board/x-go/xsort"
)

func TestSliceTransform(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		require.EqualValues(t, []interface{}{"11", "22", "33"}, Map([]interface{}{"1", "2", "3"}, func(v interface{}) interface{} {
			return strings.Repeat(v.(string), 2)
		}), "map produce another slice with all of it's value string repeated two times")
	})

	t.Run("filter", func(t *testing.T) {
		require.Equal(t, 2, len(Filter([]interface{}{1, 2, 3}, func(v interface{}) bool {
			return v.(int)%2 == 1
		})), "filter produce another slice with length 2")
		require.EqualValues(t, []interface{}{1, 3}, Filter([]interface{}{1, 2, 3}, func(v interface{}) bool {
			return v.(int)%2 == 1
		}), "filter produce another slice with two elements")
	})

	t.Run("fold", func(t *testing.T) {
		t.Run("left", func(t *testing.T) {
			require.Equal(t, -6, FoldLeft([]interface{}{1, 2, 3}, func(left, right interface{}) interface{} {
				return left.(int) - right.(int)
			}, 0), "fold left produce -6")
		})
		t.Run("right", func(t *testing.T) {
			require.Equal(t, 0, FoldRight([]interface{}{1, 2, 3}, func(left, right interface{}) interface{} {
				return left.(int) - right.(int)
			}, 6), "fold left produce 0")
		})
	})
}

func TestUniqueSlice(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		x := xsort.IntSlice(UniqueIntSlice([]int{1, 2, 3, 1, 2, 3}))
		x.Sort()
		require.EqualValues(t, []int{1, 2, 3}, x, "unique int slice produce []int{1,2,3,}")
	})
}

func TestFunc(t *testing.T) {
}
