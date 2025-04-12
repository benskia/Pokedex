package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mut      *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache returns a pointer to a fresh configuration with a concurrent reaploop.
func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache:    map[string]cacheEntry{},
		mut:      &sync.RWMutex{},
		interval: interval,
	}
	go cache.reapLoop()
	return &cache
}

// cache.Add creates a new cache entry at key with byteslice val.
func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// cache.Get returns the cache entry at key and whether it was successful.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.RLock()
	defer c.mut.RUnlock()
	cacheEntry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return cacheEntry.val, true
}

// reapLoop culls cache entries from the pokecache that have exceeded the
// config's interval. Repeated per interval.
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			c.mut.Lock()
			for key, entry := range c.cache {
				if t.Sub(entry.createdAt) > c.interval {
					delete(c.cache, key)
				}
			}
			c.mut.Unlock()
		}
		time.Sleep(c.interval)
	}
}
