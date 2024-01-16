package cache

import (
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(t time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}

	go c.PurgeLoop(t)

	return c
}

func (c *Cache) AddEntry(key string, val []byte) (cacheEntry, error) {
	if _, ok := c.cache[key]; !ok {
		c.cache[key] = cacheEntry{
			val:       val,
			createdAt: time.Now(),
		}
	}
	return c.cache[key], nil
}

func (c *Cache) GetEntry(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) PurgeLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	for range ticker.C {
		c.PurgeCache(t)
	}
}

func (c *Cache) PurgeCache(t time.Duration) {
	interval := time.Now().Add(-t)
	for key, entry := range c.cache {
		if entry.createdAt.Before(interval) {
			delete(c.cache, key)
		}
	}
}
