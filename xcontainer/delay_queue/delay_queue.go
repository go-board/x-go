package delay_queue

import (
	"math"
	"time"

	"github.com/go-board/x-go/types"
	"github.com/go-board/x-go/xcontainer/priority_queue"
)

type Task struct {
	data interface{}
	at   time.Time
}

func (t Task) Compare(o types.Comparable) types.Ordering {
	oo := o.(*Task)
	if t.at.Before(oo.at) {
		return types.OrderingLess
	}
	if t.at.Equal(oo.at) {
		return types.OrderingEqual
	}
	return types.OrderingGreater
}

func NewTask(v interface{}, at time.Time) *Task {
	return &Task{data: v, at: at}
}

func NewDelayTask(v interface{}, d time.Duration) *Task {
	return NewTask(v, time.Now().Add(d))
}

type DelayQueue struct {
	q      *priority_queue.PriorityQueue
	notify chan struct{}
}

func New() *DelayQueue {
	return &DelayQueue{
		q:      priority_queue.NewComparablePriorityQueue(false),
		notify: make(chan struct{}, 1),
	}
}

// Push new data into delay queue.
func (q *DelayQueue) Push(task *Task) {
	q.q.Push(task)
}

// Pop try to pop nearest expired data.
func (q *DelayQueue) Pop() (interface{}, bool) {
	val, _, ok := q.pop()
	return val, ok
}

func (q *DelayQueue) pop() (interface{}, time.Duration, bool) {
	task := q.q.Pop()
	if task == nil {
		return nil, time.Duration(math.MaxInt64), false
	}
	t := task.(*Task)
	now := time.Now()
	duration := t.at.Sub(now)
	if !t.at.After(now) {
		return t.data, 0, true
	}
	q.q.Push(task)
	return nil, duration, false
}

// BlockPop must get a data otherwise wait for data ready.
func (q *DelayQueue) BlockPop() interface{} {
	for {
		// todo: optimization CPU usage
		//   1. try pop, if got, return, else go to step 2
		//   2. get task awake time, and block on it, and listen on notify channel for pushing new data
		//   3. if timeout or receive notification, go to step 1
		v, ok := q.Pop()
		if ok {
			return v
		}
	}
}
