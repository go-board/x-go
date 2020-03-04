package xstrings

import (
	"testing"
)

// BenchmarkRandom/random-12         	  188780	      5898 ns/op	    1720 B/op	      99 allocs/op
// BenchmarkRandom/fast-12           	 3107966	       389 ns/op	     176 B/op	       4 allocs/op
func BenchmarkRandom(b *testing.B) {
	b.Run("random", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			Random(RandomSet, 32)
		}
	})
	b.Run("fast", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			FastRandom(32)
		}
	})
}
