package priority_queue

import (
	"testing"

	"github.com/go-board/x-go/types"

	"github.com/stretchr/testify/assert"
)

func TestStringPriorityQueue(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		q := NewStringPriorityQueue(true, "1", "2", "3")
		t.Logf("%+v\n", q.h)
		assert.Equal(t, "3", q.Pop())
	})

	t.Run("min", func(t *testing.T) {
		q := NewStringPriorityQueue(false, "1", "2", "3")
		assert.Equal(t, "1", q.Pop())
	})
}

func TestInt64PriorityQueue(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		q := NewInt64PriorityQueue(true, 1, 3, 5, 2, 4, 6)
		assert.Equal(t, q.Pop(), int64(6))
	})

	t.Run("min", func(t *testing.T) {
		q := NewInt64PriorityQueue(false, 1, 3, 5, 2, 4, 6)
		assert.Equal(t, q.Pop(), int64(1))
	})
}

type String string

func (s String) Compare(o types.Comparable) types.Ordering {
	if s < o.(String) {
		return types.OrderingLess
	} else if s > o.(String) {
		return types.OrderingGreater
	}
	return types.OrderingEqual
}

func TestComparablePriorityQueue(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		q := NewComparablePriorityQueue(true, String("1"), String("4"), String("3"), String("2"))
		assert.Equal(t, q.Pop(), String("4"))
	})

	t.Run("min", func(t *testing.T) {
		q := NewComparablePriorityQueue(false, String("1"), String("4"), String("3"), String("2"))
		assert.Equal(t, q.Pop(), String("1"))
	})
}
