package set

import (
	"github.com/google/btree"
)

type StringItem string

func (s StringItem) Less(than btree.Item) bool {
	return s < than.(StringItem)
}

type StringBtreeSet struct {
	tree *btree.BTree
}

func NewStringBtreeSet(items ...string) *StringBtreeSet {
	tree := btree.New(64)
	for _, item := range items {
		tree.ReplaceOrInsert(StringItem(item))
	}
	set := &StringBtreeSet{tree: tree}
	return set
}

func (s *StringBtreeSet) Add(items ...string) {
	for _, item := range items {
		s.tree.ReplaceOrInsert(StringItem(item))
	}
}

func (s *StringBtreeSet) Remove(items ...string) {
	for _, item := range items {
		s.tree.Delete(StringItem(item))
	}
}

func (s *StringBtreeSet) Values() []string {
	values := make([]string, 0, s.tree.Len())
	s.tree.Ascend(func(i btree.Item) bool {
		values = append(values, string(i.(StringItem)))
		return true
	})
	return values
}

func (s *StringBtreeSet) RangeAll(fn func(item string)) {
	s.tree.Ascend(func(i btree.Item) bool {
		fn(string(i.(StringItem)))
		return true
	})
}

func (s *StringBtreeSet) Contains(item string) bool {
	return s.tree.Has(StringItem(item))
}

func (s *StringBtreeSet) ContainAll(items ...string) bool {
	for _, item := range items {
		if !s.tree.Has(StringItem(item)) {
			return false
		}
	}
	return true
}

func (s *StringBtreeSet) ContainAny(items ...string) bool {
	for _, item := range items {
		if s.tree.Has(StringItem(item)) {
			return true
		}
	}
	return false
}

func (s *StringBtreeSet) Size() int {
	return s.tree.Len()
}

func (s *StringBtreeSet) Union(o *StringBtreeSet) *StringBtreeSet {
	n := NewStringBtreeSet(s.Values()...)
	o.RangeAll(func(item string) {
		n.Add(item)
	})
	return n
}

func (s *StringBtreeSet) InterSection(o *StringBtreeSet) *StringBtreeSet {
	n := NewStringBtreeSet()
	s.RangeAll(func(item string) {
		if o.ContainAll(item) {
			n.Add(item)
		}
	})
	return n
}

func (s *StringBtreeSet) Difference(o *StringBtreeSet) *StringBtreeSet {
	n := NewStringBtreeSet(s.Values()...)
	n.Remove(o.Values()...)
	return n
}

func (s *StringBtreeSet) IsSuperSet(o *StringBtreeSet) bool {
	return s.ContainAll(o.Values()...)
}

func (s *StringBtreeSet) IsSubSet(o *StringBtreeSet) bool {
	return o.ContainAll(s.Values()...)
}

func (s *StringBtreeSet) Range(begin, end string, fn func(key string) bool) {
	s.tree.AscendGreaterOrEqual(StringItem(begin), func(i btree.Item) bool {
		item := i.(StringItem)
		return string(item) < end && fn(string(item))
	})
}
