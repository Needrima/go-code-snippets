package cache

import "sync"

type Cache struct {
	mu sync.Mutex
	memory map[string]string
}

func New() *Cache {
	return &Cache{memory: make(map[string]string)}
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	c.memory[key] = value
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.memory[key]
	return value, ok
}
