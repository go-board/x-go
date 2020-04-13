package maps

import (
	"hash/crc32"
	"sync"
)

type StringKeyHashMap struct {
	shard  uint32
	mus    []sync.RWMutex
	ms     []map[string]interface{}
	hashFn func(key string) uint32
}

func NewStringKeyHashMap(shard uint32, fn func(key string) uint32) *StringKeyHashMap {
	ms := make([]map[string]interface{}, shard)
	for i := 0; i < int(shard); i++ {
		ms[i] = make(map[string]interface{})
	}
	if fn == nil {
		fn = func(key string) uint32 { return crc32.ChecksumIEEE([]byte(key)) }
	}
	return &StringKeyHashMap{
		shard:  shard,
		mus:    make([]sync.RWMutex, shard),
		ms:     ms,
		hashFn: fn,
	}
}

func (m *StringKeyHashMap) hashKey(key string) uint32 {
	return m.hashFn(key) % m.shard
}

func (m *StringKeyHashMap) Set(key string, val interface{}) {
	h := m.hashKey(key)
	m.mus[h].Lock()
	m.ms[h][key] = val
	m.mus[h].Unlock()
}

func (m *StringKeyHashMap) GetOk(key string) (interface{}, bool) {
	h := m.hashKey(key)
	m.mus[h].RLock()
	val, ok := m.ms[h][key]
	m.mus[h].RUnlock()
	return val, ok
}

func (m *StringKeyHashMap) Remove(key string) {
	h := m.hashKey(key)
	m.mus[h].Lock()
	if _, ok := m.ms[h][key]; ok {
		delete(m.ms[h], key)
	}
	m.mus[h].Unlock()
}

func (m *StringKeyHashMap) RemoveAll() {
	newM := NewStringKeyHashMap(m.shard, m.hashFn)
	*m = *newM
}
