package xatomic

import (
	"sync/atomic"
)

type AtomicString struct {
	v atomic.Value
}

func NewAtomicString(str string) *AtomicString {
	v := atomic.Value{}
	v.Store(str)
	return &AtomicString{v: v}
}

func (s *AtomicString) Load() string {
	return s.v.Load().(string)
}

func (s *AtomicString) Store(str string) {
	s.v.Store(str)
}
