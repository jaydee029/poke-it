package pokecache

import "time"

type Cache struct {
	cache map[string]returnCache
}

type returnCache struct {
	val  []byte
	time time.Time
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]returnCache),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = returnCache{
		val:  val,
		time: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	ck, ok := c.cache[key]

	return ck.val, ok
}
