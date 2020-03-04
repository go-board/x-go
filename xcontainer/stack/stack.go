package stack

type Stack interface {
	Push(x interface{})
	Pop() (x interface{}, ok bool)
	Peek() (x interface{}, ok bool)
	Size() int
}
