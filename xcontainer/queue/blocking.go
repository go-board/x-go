package queue

import (
	"time"
)

type BlockingQueue struct {
	capacity int
	ch       chan interface{}
}

func (q *BlockingQueue) Capacity() int {
	return q.capacity
}

func (q *BlockingQueue) Put(x interface{}, timeout time.Duration) bool {
	t := time.NewTimer(timeout)
	select {
	case q.ch <- x:
		return true
	case <-t.C:
		return false
	}
}

func (q *BlockingQueue) Get(timeout time.Duration) (interface{}, bool) {
	t := time.NewTimer(timeout)
	select {
	case <-t.C:
		return nil, false
	case x := <-q.ch:
		return x, true
	}
}

func NewBlockingQueue(n int) *BlockingQueue {
	return &BlockingQueue{
		capacity: n,
		ch:       make(chan interface{}, n),
	}
}
