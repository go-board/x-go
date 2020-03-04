package priority_queue

import (
	"container/heap"
)

type PriorityQueue struct {
	h heap.Interface
}

func (q *PriorityQueue) Push(x interface{}) {
	heap.Push(q.h, x)
}

func (q *PriorityQueue) Pop() interface{} {
	return heap.Pop(q.h)
}

func NewPriorityQueue(h heap.Interface) *PriorityQueue {
	return &PriorityQueue{h: h}
}
