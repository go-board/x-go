package xatomic

import (
	"sync/atomic"
)

func boolToInt(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

func truthy(n uint32) bool {
	return n&1 == 1
}

// AtomicBool is an atomic Boolean.
type AtomicBool struct{ v uint32 }

// NewAtomicBool creates a AtomicBool.
func NewAtomicBool(initial bool) *AtomicBool {
	return &AtomicBool{boolToInt(initial)}
}

// Load atomically loads the Boolean.
func (b *AtomicBool) Load() bool {
	return truthy(atomic.LoadUint32(&b.v))
}

// CAS is an atomic compare-and-swap.
func (b *AtomicBool) CAS(old, new bool) bool {
	return atomic.CompareAndSwapUint32(&b.v, boolToInt(old), boolToInt(new))
}

// Store atomically stores the passed value.
func (b *AtomicBool) Store(new bool) {
	atomic.StoreUint32(&b.v, boolToInt(new))
}

// Swap sets the given value and returns the previous value.
func (b *AtomicBool) Swap(new bool) bool {
	return truthy(atomic.SwapUint32(&b.v, boolToInt(new)))
}

// Toggle atomically negates the Boolean and returns the previous value.
func (b *AtomicBool) Toggle() bool {
	return truthy(atomic.AddUint32(&b.v, 1) - 1)
}
