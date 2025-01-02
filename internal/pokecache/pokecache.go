package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entry map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		Entry: make(map[string]cacheEntry),
		mu: sync.Mutex{},
	}

	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	newCacheEntry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entry[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entryStruct, exists := c.Entry[key]
	if !exists {
		return []byte{}, false
	}

	return entryStruct.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		<-ticker.C

		c.mu.Lock()

		for key, entry := range c.Entry {
			if time.Since(entry.createdAt) > interval {
				delete(c.Entry, key)
			}  
		}
		c.mu.Unlock()
	} 
}
