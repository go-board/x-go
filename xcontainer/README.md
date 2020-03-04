# XContainer

go作为强类型，静态类型语言，因为缺乏泛型，缺少了大量的通用集合库，因此这里旨在提供常用的集合库，减少模版代码。

## maps
##### StringKeyBtreeMap 基于B树的key为string的map
定义
```go
type Entry struct {
	key string
	val interface{}
}

type StringKeyBtreeMap struct {
	tree *btree.BTree
}
```
方法
```go
func (*Entry) Less(than btree.Item) bool
func (*StringKeyBtreeMap) Set(key string, val interface{})
func (*StringKeyBtreeMap) GetOk(key string) (interface{}, bool)
func (*StringKeyBtreeMap) Remove(key string)
func (*StringKeyBtreeMap) Exist(key string) bool
func (*StringKeyBtreeMap) Range(begin, end string, fn func(key string, val interface{}) bool)
func (*StringKeyBtreeMap) RangeAll(fn func(key string, val interface{}))
```
## priority_queue 优先级队列
定义
```go
type PriorityQueue struct {
	h heap.Interface
}
```
方法
```go
func NewPriorityQueue(h heap.Interface) *PriorityQueue
func (*PriorityQueue) Push(x interface{})
func (*PriorityQueue) Pop() interface{}
```
支持的优先级队列的元素类型有`types.Comparable`, `int64`, `string`
## queue 双向队列

## set 集合

## stack 栈
