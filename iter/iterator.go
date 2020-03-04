package iter

type Iter interface {
	Iter() Iterator
}

type Iterator interface {
	Next() (interface{}, bool)
}

type IteratorExt interface {
	Iterator
	Map(fn func(item interface{}) interface{}) *Map
}

type IntoIterator interface {
	IntoIter() Iterator
}
