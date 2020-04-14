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
	// this should not block
	select {
	case q.notify <- struct{}{}:
	default:
	}
}

// Pop try to pop nearest expired data.
func (q *DelayQueue) Pop() (interface{}, bool) {
	val, _, ok := q.popNearest()
	return val, ok
}

func (q *DelayQueue) popNearest() (interface{}, time.Duration, bool) {
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
		v, duration, ok := q.popNearest()
		if ok {
			// drain notification if possible, otherwise we may get old notification
			select {
			case <-q.notify:
			default:
			}
			return v
		}
		select {
		case <-time.NewTimer(duration).C:
		case <-q.notify:
		}
	}
}
