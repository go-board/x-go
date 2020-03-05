package xsql

import (
	"hash/crc32"
	"math/rand"
	"sync/atomic"
)

type Selector interface {
	SelectDB(n int) int
}

type SelectorFunc func(n int) int

func (fn SelectorFunc) SelectDB(n int) int { return fn(n) }

func RoundRobinSelector() SelectorFunc {
	var idx int64
	return func(n int) int {
		return int(atomic.AddInt64(&idx, 1)) % n
	}
}

func ShardSelector(key string) SelectorFunc {
	return func(n int) int {
		return int(crc32.ChecksumIEEE([]byte(key))) % n
	}
}

func RandomSelector() SelectorFunc {
	return func(n int) int {
		return rand.Intn(n)
	}
}
