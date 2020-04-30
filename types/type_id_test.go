package types

import (
	"testing"
)

func TestTypeof(t *testing.T) {
	t.Log(Typeof(0))
}

func BenchmarkTypeof(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Typeof(0)
	}
}
