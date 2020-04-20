package queue

import (
	"time"
)

type BlockQueue struct {
	capacity int
	ch       chan interface{}
}

func (q *BlockQueue) Capacity() int {
	return q.capacity
}

func (q *BlockQueue) Push(x interface{}, timeout time.Duration) bool {
	t := time.NewTimer(timeout)
	select {
	case q.ch <- x:
		return true
	case <-t.C:
		return false
	}
}

func (q *BlockQueue) Pop(timeout time.Duration) (interface{}, bool) {
	t := time.NewTimer(timeout)
	select {
	case <-t.C:
		return nil, false
	case x := <-q.ch:
		return x, true
	}
}

func (q *BlockQueue) BlockPop() (interface{}, bool) {
	val, ok := <-q.ch
	return val, ok
}

func NewBlockingQueue(n int) *BlockQueue {
	return &BlockQueue{
		capacity: n,
		ch:       make(chan interface{}, n),
	}
}
