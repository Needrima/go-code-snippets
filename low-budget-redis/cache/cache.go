package cache

import "sync"

type Cache struct {
	mu sync.Mutex
	db map[string]string
}

func New() *Cache {
	return &Cache{db: make(map[string]string)}
}

func (m *Cache) Set(key, value string) {
	m.mu.Lock()
	m.db[key] = value
	m.mu.Unlock()
}

func (m *Cache) Get(key string) (string, bool) {
	value, ok := m.db[key]
	return value, ok
}
