package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cache: map[string]cacheEntry{},
	}
	go c.reapLoop(interval)

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		time.Now(),
		val,
	}
}

func (c *Cache) Get(key string) (val []byte, exists bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, exists := c.cache[key]; exists {
		return entry.val, exists
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range c.cache {
		if time.Since(val.createdAt) > interval {
			delete(c.cache, key)
		}
	}
}
