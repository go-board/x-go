package xslice

func UniqueSlice(list []interface{}) []interface{} {
	res := make(map[interface{}]struct{})
	for _, x := range list {
		res[x] = struct{}{}
	}
	result := make([]interface{}, 0)
	for k := range res {
		result = append(result, k)
	}
	return result
}

func UniqueStringSlice(list []string) []string {
	res := make(map[string]struct{})
	for _, x := range list {
		res[x] = struct{}{}
	}
	result := make([]string, 0)
	for k := range res {
		result = append(result, k)
	}
	return result
}

func UniqueIntSlice(list []int) []int {
	res := make(map[int]struct{})
	for _, x := range list {
		res[x] = struct{}{}
	}
	result := make([]int, 0)
	for k := range res {
		result = append(result, k)
	}
	return result
}

func UniqueInt64Slice(list []int64) []int64 {
	res := make(map[int64]struct{})
	for _, x := range list {
		res[x] = struct{}{}
	}
	result := make([]int64, 0)
	for k := range res {
		result = append(result, k)
	}
	return result
}

func UniqueInt32Slice(list []int32) []int32 {
	res := make(map[int32]struct{})
	for _, x := range list {
		res[x] = struct{}{}
	}
	result := make([]int32, 0)
	for k := range res {
		result = append(result, k)
	}
	return result
}

func UniqueInt16Slice(list []int16) []int16 {
	res := make(map[int16]struct{})
	for _, x := range list {
		res[x] = struct{}{}
	}
	result := make([]int16, 0)
	for k := range res {
		result = append(result, k)
	}
	return result
}

func UniqueInt8Slice(list []int8) []int8 {
	res := make(map[int8]struct{})
	for _, x := range list {
		res[x] = struct{}{}
	}
	result := make([]int8, 0)
	for k := range res {
		result = append(result, k)
	}
	return result
}
