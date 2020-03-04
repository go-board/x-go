package queue

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/smartystreets-prototypes/go-disruptor"
	"github.com/stretchr/testify/require"
)

func TestBlockingQueue(t *testing.T) {
	t.Run("put", func(t *testing.T) {
		q := NewBlockingQueue(1)
		require.Equal(t, true, q.Put(1, time.Second), "should not block when putting first element")
		require.Equal(t, false, q.Put(2, time.Second), "should block when putting second element")
	})
	t.Run("get", func(t *testing.T) {
		q := NewBlockingQueue(1)
		require.Equal(t, true, q.Put(1, time.Second), "should not block when putting first element")
		x, ok := q.Get(time.Second)
		require.Equal(t, true, ok, "should get first element")
		require.Equal(t, 1, x, "first element should be 1")
		_, ok = q.Get(time.Second)
		require.Equal(t, false, ok, "should get nothing when get again")
	})
}

type testItem struct {
	i int
}

const RingBufferCapacity = 1 << 14 // must be a power of 2
const RingBufferMask = RingBufferCapacity - 1

// this instance will be shared among producers and consumers of this application
var ringBuffer = [RingBufferCapacity]int{}

type disruptorConsumer struct{}

func (d disruptorConsumer) Consume(lower, upper int64) {
	for sequence := lower; sequence <= upper; sequence++ {

		_ = ringBuffer[sequence&RingBufferMask] // see performance note on producer sample above

		// handle the incoming message with your application code
	}
}

func BenchmarkDisruptor(b *testing.B) {
	b.Run("put", func(b *testing.B) {
		writer, _ := disruptor.New(
			disruptor.WithCapacity(1<<14),
			disruptor.WithConsumerGroup(disruptorConsumer{}))
		b.ResetTimer()
		b.ReportAllocs()
		i := int64(0)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				reservation := writer.Reserve(1)
				ringBuffer[atomic.AddInt64(&i, 1)&RingBufferMask] = 42 // example of incoming value from a network operation such as HTTP, TCP, UDP, etc.
				writer.Commit(reservation, reservation)
			}
		})
	})
}

func BenchmarkChan(b *testing.B) {
	b.Run("get", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		ch := make(chan testItem, 10000)
		go func() {
			for {
				ch <- testItem{i: 0}
			}
		}()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				<-ch
			}
		})
	})
	b.Run("put", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		ch := make(chan testItem, 10000)
		go func() {
			for {
				<-ch
			}
		}()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ch <- testItem{i: 0}
			}
		})
	})
}
