package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntSet(t *testing.T) {
	t.Run("getset", func(t *testing.T) {
		set := IntSet{}
		require.Equal(t, true, set.Insert(1), "first putting 1 success")
		require.Equal(t, true, set.Insert(2), "first putting 2 success")
		require.Equal(t, false, set.Insert(1), "second putting 1 failed")
		require.Equal(t, true, set.Has(1), "set has 1")
		require.Equal(t, true, set.Has(2), "set dose't has 3")
	})
}

func TestStringHashSet(t *testing.T) {
	t.Run("getset", func(t *testing.T) {
		set := NewStringHashSet("1", "2", "3", "4", "5", "1")
		require.Equal(t, 5, set.Size(), "set contains 5 elements")
		require.Equal(t, true, set.Contains("2"), "set contains 2")
		require.Equal(t, false, set.Contains("6"), "set not contains 6")
	})
}

func TestStringBtreeSet(t *testing.T) {
	t.Run("getset", func(t *testing.T) {
		set := NewStringBtreeSet("1", "2", "3", "4", "5", "1")
		require.Equal(t, 5, set.Size(), "set contains 5 elements")
		require.Equal(t, true, set.Contains("2"), "set contains 2")
		require.Equal(t, false, set.Contains("6"), "set not contains 6")
	})
}
