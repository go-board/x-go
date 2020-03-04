package xos_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-board/x-go/xos"
)

func init() {
	os.Setenv("FALSE_KEY", "0")
	os.Setenv("TRUE_KEY", "1")

	os.Setenv("INT_KEY", "12345678")
	os.Setenv("INTS_KEY", "1,2,3,4,5")
	os.Setenv("BAD_INT_KEY", "asf")

	os.Setenv("STRING_KEY", "HELLO")
	os.Setenv("STRINGS_KEY", "Hello,world")
}

func TestEnvBool(t *testing.T) {
	ok, err := xos.EnvBool("FALSE_KEY")
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Error("expect false, got true")
	}

	ok, err = xos.EnvBool("TRUE_KEY")
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("expect true, got false")
	}

}

func TestEnvInt64(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		i, err := xos.EnvInt64("INT_KEY")
		require.Nil(t, err, "err must be nil")
		require.Equal(t, int64(12345678), i, "int value is 12345678")
	})
	t.Run("ints", func(t *testing.T) {
		is, err := xos.EnvInt64s("INTS_KEY")
		require.Nil(t, err, "err must be nil")
		require.Equal(t, 5, len(is), "slice length is 5")
	})
	t.Run("bad int", func(t *testing.T) {
		_, err := xos.EnvInt64("BAD_INT_KEY")
		require.NotNil(t, err, "err must be not nil")
	})
}

func TestEnvString(t *testing.T) {
}
