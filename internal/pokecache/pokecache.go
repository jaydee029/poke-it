package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]returnCache
	mux   *sync.RWMutex
}

type returnCache struct {
	val       []byte
	createdAt time.Time
}

func NewCache(duration time.Duration) Cache {
	c := Cache{
		cache: make(map[string]returnCache),
		mux:   &sync.RWMutex{},
	}
	go c.reapLoop(duration)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = returnCache{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	ck, ok := c.cache[key]

	return ck.val, ok
}

/*
	func (c *Cache) reapLoop(duration time.Duration) {
		ticker := time.NewTicker(duration)
		for range ticker.C {
			c.reapCache(duration)
		}
	}

	func (c *Cache) reapCache(duration time.Duration) {
		timeelapsed := time.Now().UTC().Add(-duration)
		for k, v := range c.cache {
			if v.at.Before(timeelapsed) {
				delete(c.cache, k)
			}
		}
	}
*/
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
