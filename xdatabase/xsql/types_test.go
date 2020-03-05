package xsql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringArray(t *testing.T) {
	arr := StringArray{"Hello", "world"}
	t.Run("value", func(t *testing.T) {
		value, err := arr.Value()
		assert.Nil(t, err, "err must be nil")
		assert.Equal(t, "Hello,world", value)
	})
}
