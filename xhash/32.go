package xhash

import (
	"hash/crc32"
)

// Sum32 run crc32.ChecksumIEEE
// returns the CRC-32 checksum of data using the IEEE polynomial.
func Sum32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

// Sum32String run crc32.ChecksumIEEE
// returns the CRC-32 checksum of data using the IEEE polynomial.
func Sum32String(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
