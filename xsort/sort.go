package xsort

import (
	"sort"

	"github.com/go-board/x-go/types"
)

// isNaN64 is a copy of math.IsNaN to avoid a dependency on the math package.
func isNaN64(f float64) bool { return f != f }

// isNaN32 is a copy of math.IsNaN to avoid a dependency on the math package.
func isNaN32(f float32) bool { return f != f }

type ByteSlice types.ByteSlice

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByteSlice) Sort()              { sort.Sort(p) }

type UintSlice types.UintSlice

func (p UintSlice) Len() int           { return len(p) }
func (p UintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p UintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p UintSlice) Sort()              { sort.Sort(p) }

type Uint64Slice types.Uint64Slice

func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Uint64Slice) Sort()              { sort.Sort(p) }

type Uint32Slice types.Uint32Slice

func (p Uint32Slice) Len() int           { return len(p) }
func (p Uint32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Uint32Slice) Sort()              { sort.Sort(p) }

type Uint16Slice types.Uint16Slice

func (p Uint16Slice) Len() int           { return len(p) }
func (p Uint16Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint16Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Uint16Slice) Sort()              { sort.Sort(p) }

type Uint8Slice types.Uint8Slice

func (p Uint8Slice) Len() int           { return len(p) }
func (p Uint8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Uint8Slice) Sort()              { sort.Sort(p) }

type IntSlice = sort.IntSlice

type Int64Slice types.Int64Slice

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Int64Slice) Sort()              { sort.Sort(p) }

type Int32Slice types.Int32Slice

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Int32Slice) Sort()              { sort.Sort(p) }

type Int16Slice types.Int16Slice

func (p Int16Slice) Len() int           { return len(p) }
func (p Int16Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int16Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Int16Slice) Sort()              { sort.Sort(p) }

type Int8Slice types.Int8Slice

func (p Int8Slice) Len() int           { return len(p) }
func (p Int8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Int8Slice) Sort()              { sort.Sort(p) }

type StringSlice types.StringSlice

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p StringSlice) Sort()              { sort.Sort(p) }

type Float64Slice types.Float64Slice

func (p Float64Slice) Len() int           { return len(p) }
func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN64(p[i]) && !isNaN64(p[j]) }
func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Float64Slice) Sort()              { sort.Sort(p) }

type Float32Slice types.Float32Slice

func (p Float32Slice) Len() int           { return len(p) }
func (p Float32Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN32(p[i]) && !isNaN32(p[j]) }
func (p Float32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Float32Slice) Sort()              { sort.Sort(p) }

type ComparableSlice types.ComparableSlice

func (p ComparableSlice) Len() int           { return len(p) }
func (p ComparableSlice) Less(i, j int) bool { return p[i].Compare(p[j]) == types.OrderingLess }
func (p ComparableSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ComparableSlice) Sort()              { sort.Sort(p) }
