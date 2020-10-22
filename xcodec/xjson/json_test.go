package xjson

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypes(t *testing.T) {
	t.Run("Int64String", func(t *testing.T) {
		x, err := json.Marshal(Int64String(12345))
		require.NoError(t, err)
		require.Equal(t, "\"12345\"", string(x))

		var id Int64String
		err = json.Unmarshal(x, &id)
		require.NoError(t, err)
		require.Equal(t, Int64String(12345), id)

		err = json.Unmarshal([]byte("123456789"), &id)
		require.NoError(t, err)
		require.Equal(t, Int64String(123456789), id)
	})
}
