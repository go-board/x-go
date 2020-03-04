package priority_queue

import (
	"container/heap"
)

type maxInt64Heap []int64

func (s maxInt64Heap) Len() int {
	return len(s)
}

func (s maxInt64Heap) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s maxInt64Heap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *maxInt64Heap) Push(x interface{}) {
	*s = append(*s, x.(int64))
}

func (s *maxInt64Heap) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	q := (*s)[0]
	*s = (*s)[1:]
	return q
}

type minInt64Heap []int64

func (s minInt64Heap) Len() int {
	return len(s)
}

func (s minInt64Heap) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s minInt64Heap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *minInt64Heap) Push(x interface{}) {
	*s = append(*s, x.(int64))
}

func (s *minInt64Heap) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	q := (*s)[0]
	*s = (*s)[1:]
	return q
}

func NewInt64PriorityQueue(max bool, items ...int64) *PriorityQueue {
	var h heap.Interface
	if max {
		t := maxInt64Heap(items)
		h = &t
	} else {
		t := minInt64Heap(items)
		h = &t
	}
	q := &PriorityQueue{h: h}
	heap.Init(q.h)
	return q
}
