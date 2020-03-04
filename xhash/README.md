# XHash

提供了默认的`Hash32`, `Hash64`, `Hash128`的实现, 以及一致性哈希，以及进程内的`string`和`[]byte`的`hash`。

## Hash32(CRC)
```go
package xhash

func Sum32(data []byte) uint32

func Sum32String(str string) uint32
```

## Hash64(XXHASH)
```go
package xhash

func Sum64(data []byte) uint64

func Sum64String(str string) uint64
```

## Hash128(MURMUR3)
```go
package xhash

func Sum128(data []byte) (uint64, uint64)

func Sum128String(str string) (uint64, uint64)
```

## 一致性哈希
```go
package ring

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int // Sorted
	hashMap  map[int]string
}

func New(replicas int, fn Hash) *Map

func (*Map) IsEmpty() bool // 判断hash环是否为空
func (*Map) Add(nodes ...string) // 加入节点
func (*Map) Remove(nodes ...string) // 删除节点
func (*Map) Get(key string) string // 根据key进行hash选取节点
```
## Runtime Hash (标准库在1.14中会实现)
```go
package runtime

func String(str string) uint64

func Bytes(data []byte) uint64
```
