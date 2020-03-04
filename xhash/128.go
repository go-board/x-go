package xhash

import (
	"github.com/spaolacci/murmur3"
)

func Sum128(data []byte) (uint64, uint64) {
	return murmur3.Sum128(data)
}

func Sum128String(str string) (uint64, uint64) {
	return murmur3.Sum128([]byte(str))
}
