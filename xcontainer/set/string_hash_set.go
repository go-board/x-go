package set

type StringHashSet map[string]struct{}

func NewStringHashSet(items ...string) StringHashSet {
	set := StringHashSet{}
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

func (s StringHashSet) Add(items ...string) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s StringHashSet) Remove(items ...string) {
	for _, item := range items {
		delete(s, item)
	}
}

func (s StringHashSet) Values() []string {
	values := make([]string, 0, len(s))
	for k := range s {
		values = append(values, k)
	}
	return values
}

func (s StringHashSet) Range(fn func(item string)) {
	for k := range s {
		fn(k)
	}
}

func (s StringHashSet) Contains(item string) bool {
	_, ok := s[item]
	return ok
}

func (s StringHashSet) ContainAll(items ...string) bool {
	for _, item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}

func (s StringHashSet) ContainAny(items ...string) bool {
	for _, item := range items {
		if _, ok := s[item]; ok {
			return true
		}
	}
	return false
}

func (s StringHashSet) Size() int {
	return len(s)
}

func (s StringHashSet) Union(o StringHashSet) StringHashSet {
	n := NewStringHashSet(s.Values()...)
	for k := range o {
		n[k] = struct{}{}
	}
	return n
}

func (s StringHashSet) InterSection(o StringHashSet) StringHashSet {
	n := NewStringHashSet()
	for k := range s {
		if _, ok := o[k]; ok {
			n[k] = struct{}{}
		}
	}
	return n
}

func (s StringHashSet) Difference(o StringHashSet) StringHashSet {
	n := NewStringHashSet(s.Values()...)
	n.Remove(o.Values()...)
	return n
}

func (s StringHashSet) IsSuperSet(o StringHashSet) bool {
	return s.ContainAll(o.Values()...)
}

func (s StringHashSet) IsSubSet(o StringHashSet) bool {
	return o.ContainAll(s.Values()...)
}
