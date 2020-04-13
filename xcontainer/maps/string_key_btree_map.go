package maps

import (
	"github.com/google/btree"
)

type entry struct {
	key string
	val interface{}
}

func (e *entry) Less(than btree.Item) bool {
	return e.key < than.(*entry).key
}

type StringKeyBtreeMap struct {
	tree *btree.BTree
}

func NewStringKeyBtreeMap() *StringKeyBtreeMap {
	return &StringKeyBtreeMap{
		tree: btree.New(64),
	}
}

func (m *StringKeyBtreeMap) Set(key string, val interface{}) {
	m.tree.ReplaceOrInsert(&entry{key: key, val: val})
}

func (m *StringKeyBtreeMap) GetOk(key string) (interface{}, bool) {
	item := m.tree.Get(&entry{key: key})
	if item == nil {
		return nil, false
	}
	return item.(*entry).val, true
}

func (m *StringKeyBtreeMap) Remove(key string) {
	m.tree.Delete(&entry{key: key})
}

func (m *StringKeyBtreeMap) RemoveAll() {
	m.tree = btree.New(64)
}

func (m *StringKeyBtreeMap) RemoveIf(fn func(key string, val interface{}) bool) {
	m.tree.Ascend(func(i btree.Item) bool {
		entry := i.(*entry)
		if fn(entry.key, entry.val) {
			m.tree.Delete(entry)
		}
		return true
	})
}

func (m *StringKeyBtreeMap) Exist(key string) bool {
	return m.tree.Has(&entry{key: key})
}

func (m *StringKeyBtreeMap) Range(begin, end string, fn func(key string, val interface{}) bool) {
	m.tree.AscendGreaterOrEqual(&entry{key: begin}, func(i btree.Item) bool {
		entry := i.(*entry)
		return entry.key <= end && fn(entry.key, entry.val)
	})
}

func (m *StringKeyBtreeMap) RangeAll(fn func(key string, val interface{})) {
	m.tree.Ascend(func(i btree.Item) bool {
		entry := i.(*entry)
		fn(entry.key, entry.val)
		return true
	})
}
