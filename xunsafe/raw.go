package xunsafe

import (
	"reflect"
	"unsafe"
)

// StringToBytes use a quick way to cast string to []byte without any alloc, in a 6000 length test,
// this is faster than `[]byte(string)` one 2200 times
// CAUTION: this is marked as unsafe so have limit, do not change the result.
func StringToBytes(str string) []byte {
	strHeader := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}

// BytesToString use a quick way to cast []byte to string without any alloc, in a 6000 length test,
// this is faster than `string(bytes)` one 2200 times
func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
