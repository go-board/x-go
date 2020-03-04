package xhash

import (
	"github.com/cespare/xxhash"
)

func Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}

func Sum64String(str string) uint64 {
	return xxhash.Sum64String(str)
}
