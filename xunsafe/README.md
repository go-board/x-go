# XUnsafe

在go里面， `[]byte` 和 `string` 的相互转换是非常频繁的，随之而来的是大量的内存分配，当长度越长时，对性能影响越大，所有提供了 `unsafe` 的转换方法。

> 当 `[]byte` 是由 `string` 通过 `unsafe` 转换而来的话，将不能进行更改操作，否则会panic，且无法恢复。

## StringToBytes

``` go
package main
import (
    "github.com/go-board/x-go/xunsafe"
)
func main()  {
    bs := xunsafe.StringToBytes("hello,world!")
    println(bs)
}
```

## BytesToString

``` go
package main
import (
    "github.com/go-board/x-go/xunsafe"
)

func main()  {
    str := xunsafe.BytesToString([]byte{'h', 'e', 'l', 'l', 'o', ',', 'w', 'o', 'r', 'l', 'd', '!'})
    println(str)
}
```
