package cache

import "sync"

type Cache struct {
	mu sync.Mutex
	memory map[string]string
}

func New() *Cache {
	return &Cache{memory: make(map[string]string)}
}

func (m *Cache) Set(key, value string) {
	m.mu.Lock()
	m.memory[key] = value
	m.mu.Unlock()
}

func (m *Cache) Get(key string) (string, bool) {
	value, ok := m.memory[key]
	return value, ok
}
