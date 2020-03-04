package xhash

import (
	"hash"
)

type Hash32 = hash.Hash32

type Hash64 = hash.Hash64

type Hash128 interface {
	hash.Hash
	Sum128() (uint64, uint64)
}
