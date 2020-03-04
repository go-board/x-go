# Types

拓展类型定义，方便函数的定义，以及自定义方法，用户可以直接使用。

## 基本类型
```go
type Interface interface{}

type Byte byte

type Int int
type Int64 int
type Int32 int32
type Int16 int16
type Int8 int8

type Uint uint
type Uint64 uint
type Uint32 uint32
type Uint16 uint16
type Uint8 uint8

type String string

type Float64 float64
type Float32 float32

type Bool bool

type Any = interface{}

type AnySlice = []Any
type InterfaceSlice = AnySlice

type ByteSlice = []byte

type IntSlice = []int
type Int64Slice = []int64
type Int32Slice = []int32
type Int16Slice = []int16
type Int8Slice = []int8

type UintSlice = []uint
type Uint64Slice = []uint64
type Uint32Slice = []uint32
type Uint16Slice = []uint16
type Uint8Slice = []uint8

type StringSlice = []string

type Float64Slice = []float64
type Float32Slice = []float32

type BoolSlice = []bool
```
## Comparable 比较的接口
> 未来`go`[proposal: spec: generic programming facilities](https://github.com/golang/go/issues/15292)引入泛型可以解决这个问题，但是很遗憾，可见的时间内不会有泛型
```go
// Ordering is the result of a comparison between two values.
type Ordering int

const (
    // 小于
    OrderingLess    Ordering = -1
    // 等于
    OrderingEqual            = 0
    // 大于
    OrderingGreater          = 1
)

type Comparable interface{ Compare(o Comparable) Ordering }
```
> 目前上面定义的基本类型已经实现了 `Comparable`
> 同时[xsort](../xsort/README.md)也使用到了`Comparable`
## ID
```go
type ID uint64
// 大端顺序的二进制结果
func (id ID) Binary() []byte
// 大端顺序的二进制结果的16进制的结果
func (id ID) Hex() string
// 大端顺序的二进制结果的base64的结果
func (id ID) Base64() string
```

## Convert 类型转换
```go
// int 数组到int64数组的转换
func IntSliceToInt64Slice(a types.IntSlice) types.Int64Slice
// int64 数组到int数组的转换
func Int64SliceToIntSlice(a types.Int64Slice) types.IntSlice
// int 数组到interface{}数组的转换
func IntSliceToAnySlice(a types.IntSlice) types.AnySlice
// int64 数组到interface{}数组的转换
func Int64SliceToAnySlice(a types.Int64Slice) types.AnySlice
// int 数组到string数组的转换
func IntSliceToStringSlice(a types.IntSlice) types.StringSlice
// int64 数组到string数组的转换
func Int64SliceToStringSlice(a types.Int64Slice) types.StringSlice
// string 数组到int数组的转换
func StringSliceToIntSlice(a types.StringSlice) (b types.IntSlice, hasErr bool)
// string 数组到int64数组的转换
func StringSliceToInt64Slice(a types.StringSlice) (b types.Int64Slice, hasErr bool)
```
