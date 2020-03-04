package xhash

import (
	"hash/crc32"
)

func Sum32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

func Sum32String(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
