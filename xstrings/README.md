# XStrings

标准库对string的操作比较有限，所有诞生了该库

## Enchanced join

``` go
package main

import (
    "fmt"

    "github.com/go-board/x-go/xcodec"
    _ "github.com/go-board/x-go/xcodec/json"
    "github.com/go-board/x-go/xstrings"
)

type intStringer int

func (i intStringer) String() string {
    return fmt.Sprintf("%d", i*i)
}

func main() {
    xstrings.JoinInt([]int{1, 2, 3}, "#") // Output: 1#2#3
    xstrings.JoinInt8([]int8{1, 2, 3}, "#") // Output: 1#2#3
    xstrings.JoinInt16([]int16{1, 2, 3}, "#") // Output: 1#2#3
    xstrings.JoinInt32([]int32{1, 2, 3}, "#") // Output: 1#2#3
    xstrings.JoinInt64([]int64{1, 2, 3}, "#") // Output: 1#2#3

    xstrings.JoinUint([]uint{1, 2, 3}, "#") // Output: 1#2#3
    xstrings.JoinUint8([]uint8{1, 2, 3}, "#") // Output: 1#2#3
    xstrings.JoinUint16([]uint16{1, 2, 3}, "#") // Output: 1#2#3
    xstrings.JoinUint32([]uint32{1, 2, 3}, "#") // Output: 1#2#3
    xstrings.JoinUint64([]uint64{1, 2, 3}, "#") // Output: 1#2#3
    xstrings.JoinStringer([]fmt.Stringer{intStringer(1), intStringer(2), intStringer(3)}, "#") // Output: 1#4#9
    xstrings.JoinAny([]interface{}{}, "#", xcodec.Get("json"))
}
```

## Random string

``` go
package main

import (
    "github.com/go-board/x-go/xstrings"
)

func main() {
    xstrings.Random(xstrings.RandomSet, 32) // 生成长度为32的随机字符串，字符来自xstrings.RandomSet
    xstrings.FastRandom(32) // 生成长度为32的随机字符串
}
```

## Enhanced split

``` go
package main

import (
    "github.com/go-board/x-go/xstrings"
)

func main() {
    xstrings.SplitInt("1$2$3", "$") // Output: [1, 2, 3], nil
    xstrings.SplitInt8("1$2$3", "$") // Output: [1, 2, 3], nil
    xstrings.SplitInt16("1$2$3", "$") // Output: [1, 2, 3], nil
    xstrings.SplitInt32("1$2$3", "$") // Output: [1, 2, 3], nil
    xstrings.SplitInt64("1$2$3", "$") // Output: [1, 2, 3], nil

    xstrings.SplitUint("1$2$3", "$") // Output: [1, 2, 3], nil
    xstrings.SplitUint8("1$2$3", "$") // Output: [1, 2, 3], nil
    xstrings.SplitUint16("1$2$3", "$") // Output: [1, 2, 3], nil
    xstrings.SplitUint32("1$2$3", "$") // Output: [1, 2, 3], nil
    xstrings.SplitUint64("1$2$3", "$") // Output: [1, 2, 3], nil
}
```

## Transform

``` go
package main

import (
    "github.com/go-board/x-go/xstrings"
)

func main() {
    xstrings.Reverse("Hello") // Output: olleH
    xstrings.CamelCase("hello") // Output: Hello
    xstrings.SnakeCase("HelloWorld") // Output: hello_world
    xstrings.Delete("Hello", "l") // Output: Heo
    xstrings.Count("Hello", "l") // Output: 2
}
```
