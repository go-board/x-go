package runtime

import (
	"math/rand"

	bytehash "github.com/go-board/x-go/xhash/internal/bytes"
)

var seed uint64

func init() {
	seed = rand.Uint64()
}

func String(str string) uint64 {
	s := bytehash.MakeSeed(seed)
	h := bytehash.New()
	h.SetSeed(s)
	h.AddString(str)
	return h.Sum64()
}

func Bytes(data []byte) uint64 {
	s := bytehash.MakeSeed(seed)
	h := bytehash.New()
	h.SetSeed(s)
	h.AddBytes(data)
	return h.Sum64()
}
