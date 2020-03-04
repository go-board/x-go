# XAtomic
标准库中提供了 `int32` , `int64` , `uint32` , `uint64` 的原子量，却没有提供 `float32` , `float64` , `string` , `bool` 的原子量，因此，在这里，提供了比标准库更为丰富的原子量

## AtomicBool

定义

``` go
package xatomic
type AtomicBool struct{ v uint32 }
```

使用

``` go
package main
import (
    "github.com/go-board/x-go/xsync/xatomic"
)

func main() {
    b := xatomic.NewAtomicBool(false)
    println(b.Load()) // false
    println(b.CAS(false, true)) // true
    b.Store(true)
    println(b.Swap(false)) // true
    println(b.Toggle()) // true
}
```

## AtomicFloat32/64

定义

``` go
package xatomic
type AtomicFloat64 struct {	v uint64 }
type AtomicFloat32 struct { v uint32 }
```

使用

``` go
package main

import (
    "github.com/go-board/x-go/xsync/xatomic"
)

func testAtomicFloat32() {
    f := xatomic.NewAtomicFloat32(0.01)
    println(f.Load()) // 0.01
    f.Store(2.001)
    println(f.Add(0.01)) // 2.011
    println(f.Sub(0.01)) // 2.001
    println(f.CAS(2.001, 3)) // 2.001
}

func testAtomicFloat64() {
    f := xatomic.NewAtomicFloat64(0.01)
    println(f.Load()) // 0.01
    f.Store(2.001)
    println(f.Add(0.01)) // 2.011
    println(f.Sub(0.01)) // 2.001
    println(f.CAS(2.001, 3))
}

func main() {
    testAtomicFloat32()
    testAtomicFloat64()
}
```

## Atomic(Int/Uint)(32/64)

定义

``` go
package xatomic
type AtomicInt32 struct{ v int32 }
type AtomicInt64 struct{ v int64 }
type AtomicUint32 struct{ v uint32 }
type AtomicUint64 struct{ v uint64 }
```

## AtomicString

定义

``` go
package xatomic

import (
    "sync/atomic"
)

type AtomicString struct {
	v atomic.Value
}
```
