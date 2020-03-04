package iter

import (
	"sync/atomic"

	"github.com/go-board/x-go/types"
)

type IntSlice types.IntSlice

func (i IntSlice) Iter() Iterator {
	ch := make(chan int, 1)
	go func(ch chan int) {
		for _, x := range i {
			ch <- x
		}
	}(ch)
	return &intSliceIterator{next: ch}
}

type intSliceIterator struct {
	next    chan int
	stopped int32
}

func (s *intSliceIterator) Next() (interface{}, bool) {
	if atomic.LoadInt32(&s.stopped) == 1 {
		return nil, false
	}
	item, ok := <-s.next
	if !ok {
		atomic.StoreInt32(&s.stopped, 1)
		return nil, false
	}
	return item, true
}
