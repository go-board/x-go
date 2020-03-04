package xmath

func CompareUint64(a, b uint64, max bool) uint64 {
	if max {
		if a > b {
			return a
		}
		return b
	} else {
		if a > b {
			return b
		}
		return a
	}
}

func CompareUint32(a, b uint32, max bool) uint32 {
	return uint32(CompareUint64(uint64(a), uint64(b), max))
}

func CompareUint16(a, b uint16, max bool) uint16 {
	return uint16(CompareUint64(uint64(a), uint64(b), max))
}

func CompareUint8(a, b uint8, max bool) uint8 {
	return uint8(CompareUint64(uint64(a), uint64(b), max))
}

func CompareUint(a, b uint, max bool) uint {
	return uint(CompareUint64(uint64(a), uint64(b), max))
}

func CompareInt64(a, b int64, max bool) int64 {
	if max {
		if a > b {
			return a
		}
		return b
	} else {
		if a > b {
			return b
		}
		return a
	}
}

func CompareInt32(a, b int32, max bool) int32 {
	return int32(CompareInt64(int64(a), int64(b), max))
}

func CompareInt16(a, b int16, max bool) int16 {
	return int16(CompareInt64(int64(a), int64(b), max))
}

func CompareInt8(a, b int8, max bool) int8 {
	return int8(CompareInt64(int64(a), int64(b), max))
}

func CompareInt(a, b int, max bool) int {
	return int(CompareInt64(int64(a), int64(b), max))
}
