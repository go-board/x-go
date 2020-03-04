package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("slice", func(t *testing.T) {
		s := NewSliceStack(10)
		s.Push(1)
		assert.Equal(t, 1, s.Size(), "stack size is 1 after put one element")
		s.Push(false)
		assert.Equal(t, 2, s.Size(), "stack size is 2 after put another element")
		peek, ok := s.Peek()
		assert.Equal(t, true, ok, "stack has peek element")
		assert.Equal(t, false, peek, "peek value is bool(false)")
		pop, ok := s.Pop()
		assert.Equal(t, true, ok, "stack has peek element so should be popped")
		assert.Equal(t, false, pop, "popped value is bool(false)")
		assert.Equal(t, 1, s.Size(), "stack size should be 1 after pop one element")
	})
}
