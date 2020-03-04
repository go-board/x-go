# XSort

标准库中只有 `int` , `string` , `float64` 的排序功能，其他的虽然可以使用，但是使用到了反射，对性能有影响，因此诞生了该库。

## sort. Interface 类型

``` go
package xsort

import (
    "github.com/go-board/x-go/types"
)

type ByteSlice types.ByteSlice
type UintSlice types.UintSlice
type Uint64Slice types.Uint64Slice
type Uint32Slice types.Uint32Slice
type Uint16Slice types.Uint16Slice
type Uint8Slice types.Uint8Slice
type IntSlice types.IntSlice
type Int64Slice types.Int64Slice
type Int32Slice types.Int32Slice
type Int16Slice types.Int16Slice
type Int8Slice types.Int8Slice
type StringSlice types.StringSlice
type Float64Slice types.Float64Slice
type Float32Slice types.Float32Slice
type ComparableSlice types.ComparableSlice
```

> 已上所有类型均实现了 `sort.Interface` 接口，可以直接进行 `sort` 操作

## sort. Search 实现

> 对所有类型的数据实现了 `sort.Search` 方法, 比如 `int` ，对应的 `search` 方法为 `func SearchInt(a []int, x int) int` 
