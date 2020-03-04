package priority_queue

import (
	"container/heap"
)

type maxStringHeap []string

func (s maxStringHeap) Len() int {
	return len(s)
}

func (s maxStringHeap) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s maxStringHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *maxStringHeap) Push(x interface{}) {
	*s = append(*s, x.(string))
}

func (s *maxStringHeap) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	q := (*s)[0]
	*s = (*s)[1:]
	return q
}

type minStringHeap []string

func (s minStringHeap) Len() int {
	return len(s)
}

func (s minStringHeap) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s minStringHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *minStringHeap) Push(x interface{}) {
	*s = append(*s, x.(string))
}

func (s *minStringHeap) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	q := (*s)[0]
	*s = (*s)[1:]
	return q
}

func NewStringPriorityQueue(max bool, items ...string) *PriorityQueue {
	var h heap.Interface
	if max {
		t := maxStringHeap(items)
		h = &t
	} else {
		t := minStringHeap(items)
		h = &t
	}
	q := &PriorityQueue{h: h}
	heap.Init(q.h)
	return q
}
