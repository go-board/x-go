package xhash

import (
	"hash/crc64"
)

var crc64Table = crc64.MakeTable(crc64.ECMA)

// Sum64 run crc64.Checksum
// returns the CRC-64 checksum of data using the ECMA polynomial.
func Sum64(data []byte) uint64 {
	return crc64.Checksum(data, crc64Table)
}

// Sum64String run crc64.Checksum
// returns the CRC-64 checksum of data using the ECMA polynomial.
func Sum64String(str string) uint64 {
	return crc64.Checksum([]byte(str), crc64Table)
}
