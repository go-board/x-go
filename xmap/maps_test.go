package xmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	t.Run("keys", func(t *testing.T) {
		x := Keys(map[string]int{"1": 2})
		require.EqualValues(t, []interface{}{"1"}, x, "")
		y := Keys(struct{ Key int }{Key: 11})
		require.EqualValues(t, []interface{}{"Key"}, y, "")
	})
	t.Run("values", func(t *testing.T) {
		x := Values(map[string]int{"1": 2})
		t.Logf("%+v\n", x)
		y := Values(struct{ Key int }{Key: 11})
		t.Logf("%+v\n", y)
	})
	t.Run("contains", func(t *testing.T) {
		require.Equal(t, true, Contains(map[string]struct{}{"1": {}, "2": {}}, "1"))
		require.Equal(t, false, Contains(map[string]struct{}{"1": {}, "2": {}}, "3"))
		require.Equal(t, true, Contains(map[byte]struct{}{'a': {}, 'c': {}}, uint8('c')))
		require.Equal(t, false, Contains(map[byte]struct{}{'a': {}, 'c': {}}, uint8('b')))
	})
}
