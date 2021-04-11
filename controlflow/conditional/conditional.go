package conditional

// Function execute left if v, else right.
func Function(v bool, left, right func()) {
	if v {
		left()
	} else {
		right()
	}
}

// String return left if v, else right.
func String(v bool, left, right string) string {
	if v {
		return left
	}
	return right
}

// String return left if v, else right.
func Bool(v bool, left, right bool) bool {
	if v {
		return left
	}
	return right
}

// Int64 return left if v, else right.
func Int64(v bool, left, right int64) int64 {
	if v {
		return left
	}
	return right
}

// Int32 return left if v, else right.
func Int32(v bool, left, right int32) int32 {
	if v {
		return left
	}
	return right
}

// Int16 return left if v, else right.
func Int16(v bool, left, right int16) int16 {
	if v {
		return left
	}
	return right
}

// Int8 return left if v, else right.
func Int8(v bool, left, right int8) int8 {
	if v {
		return left
	}
	return right
}

// Int return left if v, else right.
func Int(v bool, left, right int) int {
	if v {
		return left
	}
	return right
}

// Uint64 return left if v, else right.
func Uint64(v bool, left, right uint64) uint64 {
	if v {
		return left
	}
	return right
}

// Uint32 return left if v, else right.
func Uint32(v bool, left, right uint32) uint32 {
	if v {
		return left
	}
	return right
}

// Uint16 return left if v, else right.
func Uint16(v bool, left, right uint16) uint16 {
	if v {
		return left
	}
	return right
}

// Uint8 return left if v, else right.
func Uint8(v bool, left, right uint8) uint8 {
	if v {
		return left
	}
	return right
}

// Uint return left if v, else right.
func Uint(v bool, left, right uint) uint {
	if v {
		return left
	}
	return right
}

// Float32 return left if v, else right.
func Float32(v bool, left, right float32) float32 {
	if v {
		return left
	}
	return right
}

// Float64 return left if v, else right.
func Float64(v bool, left, right float64) float64 {
	if v {
		return left
	}
	return right
}

// Complex64 return left if v, else right.
func Complex64(v bool, left, right complex64) complex64 {
	if v {
		return left
	}
	return right
}

// Complex128 return left if v, else right.
func Complex128(v bool, left, right complex128) complex128 {
	if v {
		return left
	}
	return right
}

// Uintptr return left if v, else right.
func Uintptr(v bool, left, right uintptr) uintptr {
	if v {
		return left
	}
	return right
}

// Any return left if v, else right.
func Any(v bool, left, right interface{}) interface{} {
	if v {
		return left
	}
	return right
}
