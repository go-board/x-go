package xstrings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMutableStringToString(t *testing.T) {
	s := FromString("Hello,world")
	require.Equal(t, "Hello,world", s.ToString())
}

func TestMutableStringInsert(t *testing.T) {
	t.Run("Insert", func(t *testing.T) {
		s := FromString("Hello,world")
		s.Insert(0, '!')
		require.Equal(t, "!Hello,world", s.ToString())
	})
	t.Run("InsertString", func(t *testing.T) {
		s := FromString("Hello,")
		s.InsertString(s.Length(), "world")
		require.Equal(t, "Hello,world", s.ToString())
	})
}

func TestMutableStringRemove(t *testing.T) {
	t.Run("Remove", func(t *testing.T) {
		s := FromString("Hello,world")
		s.Remove(0)
		require.Equal(t, "ello,world", s.ToString())
		s.Remove(3)
		require.Equal(t, "ell,world", s.ToString())
	})
	t.Run("RemoveRange", func(t *testing.T) {
		s := FromString("Hello,world")
		s.RemoveRange(0, 2)
		require.Equal(t, "llo,world", s.ToString())
		s.RemoveRange(3, 3)
		require.Equal(t, "llorld", s.ToString())
	})
}

func TestMutableStringReplace(t *testing.T) {
	t.Run("Replace", func(t *testing.T) {
		s := FromString("Hello,world")
		s.Replace(0, 'A')
		require.Equal(t, "Aello,world", s.ToString())
		s.Replace(2, 'b')
		require.Equal(t, "Aeblo,world", s.ToString())
	})
	t.Run("ReplaceRange", func(t *testing.T) {
		s := FromString("Hello,world")
		s.ReplaceRange(0, []byte{'W', 'o', 'r', 'l', 'd'})
		require.Equal(t, "World,world", s.ToString())
		s.ReplaceRange(6, []byte{'h', 'e', 'l', 'l', 'o'})
		require.Equal(t, "World,hello", s.ToString())
	})
}
