package cache

import "sync"

type MemCache struct {
	mu sync.Mutex
	db map[string]string
}

func New() *MemCache {
	return &MemCache{db: make(map[string]string)}
}

func (m *MemCache) Set(key, value string) {
	m.mu.Lock()
	m.db[key] = value
	m.mu.Unlock()
}

func (m *MemCache) Get(key string) (string, bool) {
	value, ok := m.db[key]
	return value, ok
}
