package xstrings

import (
	"strconv"
	"strings"
)

func SplitInt(str string, sep string) ([]int, error) {
	items := strings.Split(str, sep)
	result := make([]int, 0, len(items))
	for _, item := range items {
		i, err := strconv.Atoi(item)
		if err != nil {
			return result, err
		}
		result = append(result, i)
	}
	return result, nil
}

func SplitInt8(str string, sep string) ([]int8, error) {
	items := strings.Split(str, sep)
	result := make([]int8, 0, len(items))
	for _, item := range items {
		i, err := strconv.Atoi(item)
		if err != nil {
			return result, err
		}
		result = append(result, int8(i))
	}
	return result, nil
}

func SplitInt16(str string, sep string) ([]int16, error) {
	items := strings.Split(str, sep)
	result := make([]int16, 0, len(items))
	for _, item := range items {
		i, err := strconv.Atoi(item)
		if err != nil {
			return result, err
		}
		result = append(result, int16(i))
	}
	return result, nil
}

func SplitInt32(str string, sep string) ([]int32, error) {
	items := strings.Split(str, sep)
	result := make([]int32, 0, len(items))
	for _, item := range items {
		i, err := strconv.Atoi(item)
		if err != nil {
			return result, err
		}
		result = append(result, int32(i))
	}
	return result, nil
}

func SplitInt64(str string, sep string) ([]int64, error) {
	items := strings.Split(str, sep)
	result := make([]int64, 0, len(items))
	for _, item := range items {
		i, err := strconv.Atoi(item)
		if err != nil {
			return result, err
		}
		result = append(result, int64(i))
	}
	return result, nil
}

func SplitUint(str string, sep string) ([]uint, error) {
	items := strings.Split(str, sep)
	result := make([]uint, 0, len(items))
	for _, item := range items {
		i, err := strconv.ParseUint(item, 10, 64)
		if err != nil {
			return result, err
		}
		result = append(result, uint(i))
	}
	return result, nil
}

func SplitUint8(str string, sep string) ([]uint8, error) {
	items := strings.Split(str, sep)
	result := make([]uint8, 0, len(items))
	for _, item := range items {
		i, err := strconv.ParseUint(item, 10, 64)
		if err != nil {
			return result, err
		}
		result = append(result, uint8(i))
	}
	return result, nil
}

func SplitUint16(str string, sep string) ([]uint16, error) {
	items := strings.Split(str, sep)
	result := make([]uint16, 0, len(items))
	for _, item := range items {
		i, err := strconv.ParseUint(item, 10, 64)
		if err != nil {
			return result, err
		}
		result = append(result, uint16(i))
	}
	return result, nil
}

func SplitUint32(str string, sep string) ([]uint32, error) {
	items := strings.Split(str, sep)
	result := make([]uint32, 0, len(items))
	for _, item := range items {
		i, err := strconv.ParseUint(item, 10, 64)
		if err != nil {
			return result, err
		}
		result = append(result, uint32(i))
	}
	return result, nil
}

func SplitUint64(str string, sep string) ([]uint64, error) {
	items := strings.Split(str, sep)
	result := make([]uint64, 0, len(items))
	for _, item := range items {
		i, err := strconv.ParseUint(item, 10, 64)
		if err != nil {
			return result, err
		}
		result = append(result, i)
	}
	return result, nil
}
