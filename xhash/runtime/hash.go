//+build !go1.14

package runtime

import (
	"github.com/go-board/x-go/xhash/internal/maphash"
)

func String(str string) uint64 {
	h := maphash.Hash{}
	h.WriteString(str)
	return h.Sum64()
}

func Bytes(data []byte) uint64 {
	h := maphash.Hash{}
	h.Write(data)
	return h.Sum64()
}
