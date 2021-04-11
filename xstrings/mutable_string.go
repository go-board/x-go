package xstrings

import (
	"io"
)

// MutableString is the mutable version of const string
type MutableString struct {
	vec []uint8
}

// FromString create MutableString from const string
func FromString(s string) *MutableString {
	vec := make([]byte, len(s))
	copy(vec, s)
	return &MutableString{vec: vec}
}

// WithCapacity create MutableString with given capacity.
func WithCapacity(cap int) *MutableString {
	return &MutableString{vec: make([]uint8, 0, cap)}
}

// ToString convert MutableString to const string
func (s *MutableString) ToString() string { return string(s.vec) }

// AppendString append string at tail of MutableString
func (s *MutableString) AppendString(str string) { s.vec = append(s.vec, str...) }

// AppendByte append a byte at tail of MutableString
func (s *MutableString) AppendByte(b uint8) { s.vec = append(s.vec, b) }

// AppendBytes append bytes at tail of MutableString
func (s *MutableString) AppendBytes(bytes []uint8) {
	s.vec = append(s.vec, bytes...)
}

func (s *MutableString) AsBytes() []uint8 {
	b := make([]uint8, 0, len(s.vec))
	copy(b, s.vec)
	return b
}

func (s *MutableString) Insert(idx int, b uint8) {
	if idx < 0 || idx > len(s.vec) {
		panic("")
	}
	if cap(s.vec) > len(s.vec) {
		s.vec = s.vec[:len(s.vec)+1]
		copy(s.vec[idx+1:], s.vec[idx:])
		s.vec[idx] = b
		return
	}
	newVec := make([]uint8, len(s.vec)+1)
	copy(newVec, s.vec[:idx])
	newVec[idx] = b
	copy(newVec[idx+1:], s.vec[idx:])
	s.vec = newVec
}

func (s *MutableString) InsertString(idx int, str string) {
	if idx < 0 || idx > len(s.vec) {
		panic("")
	}
	if cap(s.vec) >= len(s.vec)+len(str) {
		s.vec = s.vec[:len(s.vec)+len(str)]
		copy(s.vec[idx+len(str):], s.vec[idx:])
		copy(s.vec[idx:], str)
		return
	}
	newVec := make([]uint8, len(s.vec)+len(str))
	copy(newVec, s.vec[:idx])
	copy(newVec[idx:], str)
	copy(newVec[idx+len(str):], s.vec[idx:])
	s.vec = newVec
}

func (s *MutableString) Remove(idx int) {
	if idx < 0 || idx >= len(s.vec) {
		panic("err: index ouf of range")
	}
	copy(s.vec[idx:], s.vec[idx+1:])
	s.vec = s.vec[:len(s.vec)-1]
}

func (s *MutableString) RemoveRange(start, length int) {
	if start < 0 || length <= 0 || start > len(s.vec) || start+length > len(s.vec) {
		panic("err: index ouf of range")
	}
	copy(s.vec[start:], s.vec[start+length:])
	s.vec = s.vec[:len(s.vec)-length]
}

func (s *MutableString) Replace(idx int, b uint8) {
	if idx < 0 || idx > len(s.vec) {
		panic("err: index ouf of range")
	}
	s.vec[idx] = b
}

func (s *MutableString) ReplaceRange(idx int, bytes []uint8) {
	if idx < 0 || idx > len(s.vec) || idx+len(bytes) > len(s.vec) {
		panic("err: index ouf of range")
	}
	copy(s.vec[idx:], bytes)
}

func (s *MutableString) Length() int { return len(s.vec) }

func (s *MutableString) Capacity() int { return cap(s.vec) }

func (s *MutableString) First() (uint8, bool) {
	if len(s.vec) > 0 {
		return s.vec[0], true
	}
	return 0, false
}

func (s *MutableString) Last() (uint8, bool) {
	if len(s.vec) > 0 {
		return s.vec[len(s.vec)-1], true
	}
	return 0, false
}

func (s *MutableString) Read(p []byte) error {
	if len(p) > len(s.vec) {
		copy(p, s.vec)
		s.vec = s.vec[:0]
		return io.ErrUnexpectedEOF
	}
	copy(p, s.vec)
	s.vec = s.vec[len(p):]
	return nil
}

func (s *MutableString) Write(p []byte) (n int, err error) {
	s.AppendBytes(p)
	return len(p), nil
}
