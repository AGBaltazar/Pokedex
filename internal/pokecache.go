package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct{
	entry map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration

}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache{
	newCache := Cache{
    	entry: make(map[string]cacheEntry),
		interval: interval,
	}
	go newCache.reapLoop()

	return newCache
}
func (c *Cache) addCache(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) getCache(key string) ([]byte, bool) {
	// Retrieve a value from the map based on the given key and return both the value and a boolean indicating whether the key was found
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entry[key]
	if exists {
		return entry.val, true
	} else {
		fmt.Println("Cache not found")
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	newTicker := time.NewTicker(c.interval)
	
	for range newTicker.C {
		c.mu.Lock()
		for i := range c.entry {
			value := c.entry[i]
			if time.Since(value.createdAt) >= c.interval {
				delete(c.entry, i)
			}
		}
		c.mu.Unlock()
	}
}