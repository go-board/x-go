package maps

import (
	"github.com/google/btree"
)

type Entry struct {
	key string
	val interface{}
}

func (e *Entry) Less(than btree.Item) bool {
	return e.key < than.(*Entry).key
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
	m.tree.ReplaceOrInsert(&Entry{
		key: key,
		val: val,
	})
}

func (m *StringKeyBtreeMap) GetOk(key string) (interface{}, bool) {
	item := m.tree.Get(&Entry{key: key})
	if item == nil {
		return nil, false
	}
	return item.(*Entry).val, true
}

func (m *StringKeyBtreeMap) Remove(key string) {
	m.tree.Delete(&Entry{key: key})
}

func (m *StringKeyBtreeMap) Exist(key string) bool {
	return m.tree.Has(&Entry{key: key})
}

func (m *StringKeyBtreeMap) Range(begin, end string, fn func(key string, val interface{}) bool) {
	m.tree.AscendGreaterOrEqual(&Entry{key: begin}, func(i btree.Item) bool {
		entry := i.(*Entry)
		return entry.key <= end && fn(entry.key, entry.val)
	})
}

func (m *StringKeyBtreeMap) RangeAll(fn func(key string, val interface{})) {
	m.tree.Ascend(func(i btree.Item) bool {
		entry := i.(*Entry)
		fn(entry.key, entry.val)
		return true
	})
}
