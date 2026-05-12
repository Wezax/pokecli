package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	lock sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache{
	cache := Cache{
		entries: map[string]cacheEntry{},
		lock: sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte){
	c.lock.Lock()
	defer c.lock.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string)([]byte, bool){
	c.lock.Lock()
	defer c.lock.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration){
	for range time.NewTicker(interval).C {
		for entry := range c.entries {
			if c.entries[entry].createdAt.Compare(time.Now().Add(interval)) == -1 {
				c.lock.Lock()
				delete(c.entries, entry)
				c.lock.Unlock()
			}
		}
	}
}