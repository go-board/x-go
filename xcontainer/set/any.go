package set

type AnyHashSet map[interface{}]struct{}

func NewAnyHashSet(items ...interface{}) AnyHashSet {
	set := AnyHashSet{}
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

func (s AnyHashSet) Add(items ...interface{}) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s AnyHashSet) Remove(items ...interface{}) {
	for _, item := range items {
		delete(s, item)
	}
}

func (s AnyHashSet) Values() []interface{} {
	values := make([]interface{}, 0, len(s))
	for k := range s {
		values = append(values, k)
	}
	return values
}

func (s AnyHashSet) Range(fn func(item interface{})) {
	for k := range s {
		fn(k)
	}
}

func (s AnyHashSet) Contains(item string) bool {
	_, ok := s[item]
	return ok
}

func (s AnyHashSet) ContainAll(items ...interface{}) bool {
	for _, item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}

func (s AnyHashSet) ContainAny(items ...string) bool {
	for _, item := range items {
		if _, ok := s[item]; ok {
			return true
		}
	}
	return false
}

func (s AnyHashSet) Size() int {
	return len(s)
}

func (s AnyHashSet) Union(o AnyHashSet) AnyHashSet {
	n := NewAnyHashSet(s.Values()...)
	for k := range o {
		n[k] = struct{}{}
	}
	return n
}

func (s AnyHashSet) InterSection(o AnyHashSet) AnyHashSet {
	n := NewAnyHashSet()
	for k := range s {
		if _, ok := o[k]; ok {
			n[k] = struct{}{}
		}
	}
	return n
}

func (s AnyHashSet) Difference(o AnyHashSet) AnyHashSet {
	n := NewAnyHashSet(s.Values()...)
	n.Remove(o.Values()...)
	return n
}

func (s AnyHashSet) IsSuperSet(o AnyHashSet) bool {
	return s.ContainAll(o.Values()...)
}

func (s AnyHashSet) IsSubSet(o AnyHashSet) bool {
	return o.ContainAll(s.Values()...)
}
