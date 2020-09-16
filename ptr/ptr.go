package ptr

// Bool returns bool pointer that point to v.
func Bool(v bool) *bool { return &v }

// String returns string pointer that point to v.
func String(v string) *string { return &v }

// Int8 returns int8 pointer that point to v.
func Int8(v int8) *int8 { return &v }

// Int16 returns int16 pointer that point to v.
func Int16(v int16) *int16 { return &v }

// Int32 returns int32 pointer that point to v.
func Int32(v int32) *int32 { return &v }

// Int64 returns int64 pointer that point to v.
func Int64(v int64) *int64 { return &v }

// Int returns int pointer that point to v.
func Int(v int) *int { return &v }

// Uint8 returns uint8 pointer that point to v.
func Uint8(v uint8) *uint8 { return &v }

// Uint16 returns uint16 pointer that point to v.
func Uint16(v uint16) *uint16 { return &v }

// Uint32 returns uint32 pointer that point to v.
func Uint32(v uint32) *uint32 { return &v }

// Uint64 returns uint64 pointer that point to v.
func Uint64(v uint64) *uint64 { return &v }

// Uint returns uint pointer that point to v.
func Uint(v uint) *uint { return &v }

// Float32 returns float32 pointer that point to v.
func Float32(v float32) *float32 { return &v }

// Float64 returns float64 pointer that point to v.
func Float64(v float64) *float64 { return &v }

// Complex64 returns complex64 pointer that point to v.
func Complex64(v complex64) *complex64 { return &v }

// Complex128 returns complex128 pointer that point to v.
func Complex128(v complex128) *complex128 { return &v }

// Uintptr returns uintptr pointer that point to v.
func Uintptr(v uintptr) *uintptr { return &v }
