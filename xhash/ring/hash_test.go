package ring

import (
	"strconv"
	"testing"
)

func TestGet(t *testing.T) {
	m := New(100, nil)
	m.Add("1", "2", "3", "4", "5")
	counts := make(map[string]int, 5)
	for i := 0; i < 1000000; i++ {
		counts[m.Get(strconv.Itoa(i))]++
	}
	t.Logf("%+v\n", counts)
}

func BenchmarkGet(b *testing.B) {
	m := New(100, nil)
	m.Add("1", "2", "3", "4", "5")
	b.ResetTimer()
	b.ReportAllocs()
	//b.ReportMetric(0, "op")
	for i := 0; i < b.N; i++ {
		m.Get(strconv.Itoa(i))
	}
}
