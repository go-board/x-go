package xatomic

import (
	"math"
	"sync/atomic"
)

// AtomicFloat64 is an atomic wrapper around float64.
type AtomicFloat64 struct {
	v uint64
}

// NewAtomicFloat64 creates a AtomicFloat64.
func NewAtomicFloat64(f float64) *AtomicFloat64 {
	return &AtomicFloat64{math.Float64bits(f)}
}

// Load atomically loads the wrapped value.
func (f *AtomicFloat64) Load() float64 {
	return math.Float64frombits(atomic.LoadUint64(&f.v))
}

// Store atomically stores the passed value.
func (f *AtomicFloat64) Store(s float64) {
	atomic.StoreUint64(&f.v, math.Float64bits(s))
}

// Add atomically adds to the wrapped float64 and returns the new value.
func (f *AtomicFloat64) Add(s float64) float64 {
	for {
		old := f.Load()
		new := old + s
		if f.CAS(old, new) {
			return new
		}
	}
}

// Sub atomically subtracts from the wrapped float64 and returns the new value.
func (f *AtomicFloat64) Sub(s float64) float64 {
	return f.Add(-s)
}

// CAS is an atomic compare-and-swap.
func (f *AtomicFloat64) CAS(old, new float64) bool {
	return atomic.CompareAndSwapUint64(&f.v, math.Float64bits(old), math.Float64bits(new))
}

// AtomicFloat32 is an atomic wrapper around float32.
type AtomicFloat32 struct {
	v uint32
}

// NewAtomicFloat32 creates a AtomicFloat32.
func NewAtomicFloat32(f float32) *AtomicFloat32 {
	return &AtomicFloat32{math.Float32bits(f)}
}

// Load atomically loads the wrapped value.
func (f *AtomicFloat32) Load() float32 {
	return math.Float32frombits(atomic.LoadUint32(&f.v))
}

// Store atomically stores the passed value.
func (f *AtomicFloat32) Store(s float32) {
	atomic.StoreUint32(&f.v, math.Float32bits(s))
}

// Add atomically adds to the wrapped float32 and returns the new value.
func (f *AtomicFloat32) Add(s float32) float32 {
	for {
		old := f.Load()
		new := old + s
		if f.CAS(old, new) {
			return new
		}
	}
}

// Sub atomically subtracts from the wrapped float32 and returns the new value.
func (f *AtomicFloat32) Sub(s float32) float32 {
	return f.Add(-s)
}

// CAS is an atomic compare-and-swap.
func (f *AtomicFloat32) CAS(old, new float32) bool {
	return atomic.CompareAndSwapUint32(&f.v, math.Float32bits(old), math.Float32bits(new))
}
