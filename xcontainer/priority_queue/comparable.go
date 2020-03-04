package priority_queue

import (
	"container/heap"

	"github.com/go-board/x-go/types"

	"github.com/go-board/x-go/xsort"
)

type maxComparableHeap struct{ xsort.ComparableSlice }

func (m *maxComparableHeap) Less(i, j int) bool {
	return !m.ComparableSlice.Less(i, j)
}

func (m *maxComparableHeap) Push(x interface{}) {
	m.ComparableSlice = append(m.ComparableSlice, x.(types.Comparable))
}

func (m *maxComparableHeap) Pop() interface{} {
	if m.Len() == 0 {
		return nil
	}
	p := (m.ComparableSlice)[len(m.ComparableSlice)-1]
	m.ComparableSlice = m.ComparableSlice[:len(m.ComparableSlice)-1]
	return p
}

type minComparableHeap struct{ xsort.ComparableSlice }

func (m *minComparableHeap) Push(x interface{}) {
	m.ComparableSlice = append(m.ComparableSlice, x.(types.Comparable))
}

func (m *minComparableHeap) Pop() interface{} {
	if m.Len() == 0 {
		return nil
	}
	p := (m.ComparableSlice)[len(m.ComparableSlice)-1]
	m.ComparableSlice = m.ComparableSlice[:len(m.ComparableSlice)-1]
	return p
}

func NewComparablePriorityQueue(max bool, items ...types.Comparable) *PriorityQueue {
	var h heap.Interface
	if max {
		h = &maxComparableHeap{items}
	} else {
		h = &minComparableHeap{items}
	}
	return NewPriorityQueue(h)
}
