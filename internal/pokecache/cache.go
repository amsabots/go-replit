package pokecache

import (
	"sync"
	"time"
)


type Cache struct {
	cache map[string]cacheEntry
	mu *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}


func NewCache(dur time.Duration) *Cache{
	cacheMap := make(map[string]cacheEntry)
	return &Cache{
       cache: cacheMap,
	   mu: &sync.Mutex{},

	}
}

func (c *Cache) Add(key string, value []byte)

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration){
  ticker := time.NewTicker(interval)
  for range ticker.C {
    c.reap(time.Now().UTC(), interval)
  }
}

func (c *Cache) reap(now time.Time, last time.Duration){
	 c.mu.Lock();
	 defer c.mu.Unlock()

	 for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)){delete(c.cache, k)}
	 }
}

