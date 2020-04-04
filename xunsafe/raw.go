package xunsafe

import (
	"reflect"
	"unsafe"
)

// StringToBytes use a quick way to cast string to []byte without any alloc, in a 6000 length test,
// this is faster than `[]byte(string)` one 2200 times
// CAUTION: this is marked as unsafe so have limit, do not change the result.
// Generated asm code is:
// "".s2b STEXT nosplit size=88 args=0x28 locals=0x20
// 0x0000 00000 (main.go:33) TEXT "".s2b(SB), NOSPLIT|ABIInternal, $32-40
// 0x0000 00000 (main.go:33) SUBQ $32, SP
// 0x0004 00004 (main.go:33) MOVQ BP, 24(SP)
// 0x0009 00009 (main.go:33) LEAQ 24(SP), BP
// 0x000e 00014 (main.go:33) FUNCDATA $0, gclocals·9fad110d66c97cf0b58d28cccea80b12(SB)
// 0x000e 00014 (main.go:33) FUNCDATA $1, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
// 0x000e 00014 (main.go:33) FUNCDATA $3, gclocals·ebb0e8ce1793da18f0378b883cb3e122(SB)
// 0x000e 00014 (main.go:33) FUNCDATA $4, "".s2b.stkobj(SB)
// 0x000e 00014 (main.go:35) PCDATA $2, $0
// 0x000e 00014 (main.go:35) PCDATA $0, $0
// 0x000e 00014 (main.go:35) XORPS X0, X0
// 0x0011 00017 (main.go:35) MOVUPS X0, "".bh(SP)
// 0x0015 00021 (main.go:35) MOVQ $0, "".bh+16(SP)
// 0x001e 00030 (main.go:36) MOVQ "".s+40(SP), AX
// 0x0023 00035 (main.go:36) MOVQ AX, "".bh(SP)
// 0x0027 00039 (main.go:37) MOVQ "".s+48(SP), AX
// 0x002c 00044 (main.go:37) MOVQ AX, "".bh+8(SP)
// 0x0031 00049 (main.go:38) PCDATA $0, $1
// 0x0031 00049 (main.go:38) MOVQ "".s+48(SP), CX
// 0x0036 00054 (main.go:38) MOVQ CX, "".bh+16(SP)
// 0x003b 00059 (main.go:40) PCDATA $2, $1
// 0x003b 00059 (main.go:40) MOVQ "".bh(SP), DX
// 0x003f 00063 (main.go:40) PCDATA $2, $0
// 0x003f 00063 (main.go:40) PCDATA $0, $2
// 0x003f 00063 (main.go:40) MOVQ DX, "".~r1+56(SP)
// 0x0044 00068 (main.go:40) MOVQ AX, "".~r1+64(SP)
// 0x0049 00073 (main.go:40) MOVQ CX, "".~r1+72(SP)
// 0x004e 00078 (main.go:40) MOVQ 24(SP), BP
// 0x0053 00083 (main.go:40) ADDQ $32, SP
// 0x0057 00087 (main.go:40) RET
func StringToBytes(str string) []byte {
	strHeader := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}

// StringToBytesFast is a faster version of StringToBytes.
// return stack space to store the final value, without the allocation of a temporary struct.
// Generated asm code is:
// "".s2bFast STEXT nosplit size=43 args=0x28 locals=0x0
// 0x0000 00000 (main.go:43) TEXT "".s2bV1(SB), NOSPLIT|ABIInternal, $0-40
// 0x0000 00000 (main.go:43) FUNCDATA $0, gclocals·39d1b96ca581879f548ad2c8aeb3a5fe(SB)
// 0x0000 00000 (main.go:43) FUNCDATA $1, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
// 0x0000 00000 (main.go:43) FUNCDATA $3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 0x0000 00000 (main.go:43) FUNCDATA $4, "".s2bV1.stkobj(SB)
// 0x0000 00000 (main.go:43) PCDATA $2, $0
// 0x0000 00000 (main.go:43) PCDATA $0, $1
// 0x0000 00000 (main.go:43) MOVQ $0, "".b+24(SP)
// 0x0009 00009 (main.go:43) XORPS X0, X0
// 0x000c 00012 (main.go:43) MOVUPS X0, "".b+32(SP)
// 0x0011 00017 (main.go:45) MOVQ "".s+16(SP), AX
// 0x0016 00022 (main.go:45) PCDATA $0, $2
// 0x0016 00022 (main.go:45) MOVQ "".s+8(SP), CX
// 0x001b 00027 (main.go:46) MOVQ CX, "".b+24(SP)
// 0x0020 00032 (main.go:47) MOVQ AX, "".b+32(SP)
// 0x0025 00037 (main.go:48) MOVQ AX, "".b+40(SP)
// 0x002a 00042 (main.go:49) RET
func StringToBytesFast(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

// BytesToString use a quick way to cast []byte to string without any alloc, in a 6000 length test,
// this is faster than `string(bytes)` one 2200 times
func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
