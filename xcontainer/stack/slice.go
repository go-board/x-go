package stack

type sliceStack []interface{}

func (s *sliceStack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *sliceStack) Pop() (x interface{}, ok bool) {
	if (*s).Size() == 0 {
		return nil, false
	}
	p := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return p, true
}

func (s sliceStack) Peek() (x interface{}, ok bool) {
	if s.Size() == 0 {
		return nil, false
	}
	return s[len(s)-1], true
}

func (s sliceStack) Size() int {
	return len(s)
}

func NewSliceStack(capacity int) *sliceStack {
	stack := sliceStack(make([]interface{}, 0, capacity))
	return &stack
}
