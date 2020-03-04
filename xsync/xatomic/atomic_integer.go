package xatomic

import (
	"sync/atomic"
)

// AtomicInt32 is an atomic wrapper around an int32.
type AtomicInt32 struct{ v int32 }

// NewAtomicInt32 creates an AtomicInt32.
func NewAtomicInt32(i int32) *AtomicInt32 {
	return &AtomicInt32{i}
}

// Load atomically loads the wrapped value.
func (i *AtomicInt32) Load() int32 {
	return atomic.LoadInt32(&i.v)
}

// Add atomically adds to the wrapped int32 and returns the new value.
func (i *AtomicInt32) Add(n int32) int32 {
	return atomic.AddInt32(&i.v, n)
}

// Sub atomically subtracts from the wrapped int32 and returns the new value.
func (i *AtomicInt32) Sub(n int32) int32 {
	return atomic.AddInt32(&i.v, -n)
}

// Inc atomically increments the wrapped int32 and returns the new value.
func (i *AtomicInt32) Inc() int32 {
	return i.Add(1)
}

// Dec atomically decrements the wrapped int32 and returns the new value.
func (i *AtomicInt32) Dec() int32 {
	return i.Sub(1)
}

// CAS is an atomic compare-and-swap.
func (i *AtomicInt32) CAS(old, new int32) bool {
	return atomic.CompareAndSwapInt32(&i.v, old, new)
}

// Store atomically stores the passed value.
func (i *AtomicInt32) Store(n int32) {
	atomic.StoreInt32(&i.v, n)
}

// Swap atomically swaps the wrapped int32 and returns the old value.
func (i *AtomicInt32) Swap(n int32) int32 {
	return atomic.SwapInt32(&i.v, n)
}

// AtomicInt64 is an atomic wrapper around an int64.
type AtomicInt64 struct{ v int64 }

// NewAtomicInt64 creates an AtomicInt64.
func NewAtomicInt64(i int64) *AtomicInt64 {
	return &AtomicInt64{i}
}

// Load atomically loads the wrapped value.
func (i *AtomicInt64) Load() int64 {
	return atomic.LoadInt64(&i.v)
}

// Add atomically adds to the wrapped int64 and returns the new value.
func (i *AtomicInt64) Add(n int64) int64 {
	return atomic.AddInt64(&i.v, n)
}

// Sub atomically subtracts from the wrapped int64 and returns the new value.
func (i *AtomicInt64) Sub(n int64) int64 {
	return atomic.AddInt64(&i.v, -n)
}

// Inc atomically increments the wrapped int64 and returns the new value.
func (i *AtomicInt64) Inc() int64 {
	return i.Add(1)
}

// Dec atomically decrements the wrapped int64 and returns the new value.
func (i *AtomicInt64) Dec() int64 {
	return i.Sub(1)
}

// CAS is an atomic compare-and-swap.
func (i *AtomicInt64) CAS(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&i.v, old, new)
}

// Store atomically stores the passed value.
func (i *AtomicInt64) Store(n int64) {
	atomic.StoreInt64(&i.v, n)
}

// Swap atomically swaps the wrapped int64 and returns the old value.
func (i *AtomicInt64) Swap(n int64) int64 {
	return atomic.SwapInt64(&i.v, n)
}

// AtomicUint32 is an atomic wrapper around an uint32.
type AtomicUint32 struct{ v uint32 }

// NewAtomicUint32 creates a AtomicUint32.
func NewAtomicUint32(i uint32) *AtomicUint32 {
	return &AtomicUint32{i}
}

// Load atomically loads the wrapped value.
func (i *AtomicUint32) Load() uint32 {
	return atomic.LoadUint32(&i.v)
}

// Add atomically adds to the wrapped uint32 and returns the new value.
func (i *AtomicUint32) Add(n uint32) uint32 {
	return atomic.AddUint32(&i.v, n)
}

// Sub atomically subtracts from the wrapped uint32 and returns the new value.
func (i *AtomicUint32) Sub(n uint32) uint32 {
	return atomic.AddUint32(&i.v, ^(n - 1))
}

// Inc atomically increments the wrapped uint32 and returns the new value.
func (i *AtomicUint32) Inc() uint32 {
	return i.Add(1)
}

// Dec atomically decrements the wrapped int32 and returns the new value.
func (i *AtomicUint32) Dec() uint32 {
	return i.Sub(1)
}

// CAS is an atomic compare-and-swap.
func (i *AtomicUint32) CAS(old, new uint32) bool {
	return atomic.CompareAndSwapUint32(&i.v, old, new)
}

// Store atomically stores the passed value.
func (i *AtomicUint32) Store(n uint32) {
	atomic.StoreUint32(&i.v, n)
}

// Swap atomically swaps the wrapped uint32 and returns the old value.
func (i *AtomicUint32) Swap(n uint32) uint32 {
	return atomic.SwapUint32(&i.v, n)
}

// AtomicUint64 is an atomic wrapper around a uint64.
type AtomicUint64 struct{ v uint64 }

// NewAtomicUint64 creates a AtomicUint64.
func NewAtomicUint64(i uint64) *AtomicUint64 {
	return &AtomicUint64{i}
}

// Load atomically loads the wrapped value.
func (i *AtomicUint64) Load() uint64 {
	return atomic.LoadUint64(&i.v)
}

// Add atomically adds to the wrapped uint64 and returns the new value.
func (i *AtomicUint64) Add(n uint64) uint64 {
	return atomic.AddUint64(&i.v, n)
}

// Sub atomically subtracts from the wrapped uint64 and returns the new value.
func (i *AtomicUint64) Sub(n uint64) uint64 {
	return atomic.AddUint64(&i.v, ^(n - 1))
}

// Inc atomically increments the wrapped uint64 and returns the new value.
func (i *AtomicUint64) Inc() uint64 {
	return i.Add(1)
}

// Dec atomically decrements the wrapped uint64 and returns the new value.
func (i *AtomicUint64) Dec() uint64 {
	return i.Sub(1)
}

// CAS is an atomic compare-and-swap.
func (i *AtomicUint64) CAS(old, new uint64) bool {
	return atomic.CompareAndSwapUint64(&i.v, old, new)
}

// Store atomically stores the passed value.
func (i *AtomicUint64) Store(n uint64) {
	atomic.StoreUint64(&i.v, n)
}

// Swap atomically swaps the wrapped uint64 and returns the old value.
func (i *AtomicUint64) Swap(n uint64) uint64 {
	return atomic.SwapUint64(&i.v, n)
}
