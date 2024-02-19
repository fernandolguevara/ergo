package lib

import (
	"sync"
)

type Map[K comparable, V any] struct {
	sync.RWMutex
	m map[K]V
}

func (m *Map[K, V]) Load(key K) (V, bool) {
	m.RLock()
	v, found := m.m[key]
	m.RUnlock()
	return v, found
}

func (m *Map[K, V]) Store(key K, value V) {
	m.Lock()
	if m.m == nil {
		m.m = make(map[K]V)
	}
	m.m[key] = value
	m.Unlock()
}

func (m *Map[K, V]) Delete(key K) {
	m.Lock()
	delete(m.m, key)
	m.Unlock()
}

func (m *Map[K, V]) Range(f func(k K, v V) bool) {
	m.RLock()
	for mk, mv := range m.m {
		if f(mk, mv) == false {
			break
		}
	}
	m.RUnlock()
}
