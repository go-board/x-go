package queue

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBlockingQueue(t *testing.T) {
	t.Run("put", func(t *testing.T) {
		q := NewBlockingQueue(1)
		require.Equal(t, true, q.Push(1, time.Second), "should not block when putting first element")
		require.Equal(t, false, q.Push(2, time.Second), "should block when putting second element")
	})
	t.Run("get", func(t *testing.T) {
		q := NewBlockingQueue(1)
		require.Equal(t, true, q.Push(1, time.Second), "should not block when putting first element")
		x, ok := q.Pop(time.Second)
		require.Equal(t, true, ok, "should get first element")
		require.Equal(t, 1, x, "first element should be 1")
		_, ok = q.Pop(time.Second)
		require.Equal(t, false, ok, "should get nothing when get again")
	})
}
