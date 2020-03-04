package types

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

type ComparableSlice = []Comparable
