package iter

type Map struct {
	iter Iterator
	fn   func(item interface{}) interface{}
}

func (m *Map) Next() (interface{}, bool) {
	item, ok := m.iter.Next()
	if !ok {
		return nil, false
	}
	return m.fn(item), ok
}
