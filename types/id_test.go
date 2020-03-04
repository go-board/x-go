package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestID(t *testing.T) {
	id := ID(123456789)
	data, err := json.Marshal(id)
	require.Nil(t, err, "marshal must success")
	newId := new(ID)
	err = json.Unmarshal(data, newId)
	require.Nil(t, err, "unmarshal must success")
	require.Equal(t, ID(123456789), *newId, "data must be 123456789")
}

func TestIDMarshal(t *testing.T) {
	id := ID(1234567890987654321)
	rawID, err := json.Marshal(id)
	require.Nil(t, err, "marshal must success")
	t.Logf("raw id is %s", rawID)
}

func BenchmarkIDMarshal(b *testing.B) {
	id := ID(1234567890987654321)
	rawID := []byte("V55K6jgLnPngp")
	b.Run("marshal", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			json.Marshal(id)
		}
	})
	b.Run("unmarshal", func(b *testing.B) {
		idPtr := new(ID)
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			json.Unmarshal(rawID, idPtr)
		}
	})
}
