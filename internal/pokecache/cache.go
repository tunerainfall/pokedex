package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Data      []byte
}

type Cache struct {
	interval time.Duration
	cache    map[string]cacheEntry
	mux      *sync.Mutex
}

func NewCache(interval time.Duration) (Cache, error) {
	c := Cache{
		interval: interval,
		cache:    map[string]cacheEntry{},
		mux:      &sync.Mutex{},
	}

	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			c.reapLoop()
		}
	}()

	return c, nil
}

func (c Cache) reapLoop() {
	// see if entry has expired and remove entry
	now := time.Now()
	c.mux.Lock()
	defer c.mux.Unlock()

	for k, entry := range c.cache {
		if entry.CreatedAt.Add(c.interval).Before(now) {
			fmt.Printf("cleared cache for %s\n", k)
			delete(c.cache, k)
		}
	}
}

func (c Cache) Add(key string, value []byte) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = cacheEntry{
		CreatedAt: time.Now(),
		Data:      value,
	}

	return nil
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	val, ok := c.cache[key]
	return val.Data, ok
}
