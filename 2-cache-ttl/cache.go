package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value interface{}, ttl time.Duration)
	Delete(key string)
	Get(key string) (interface{}, error)
}

type entry struct {
	value interface{}
	timer *time.Timer
}

type inMemoryCache struct {
	data map[string]*entry
	mu   *sync.RWMutex
}

func New() Cache {
	cache := &inMemoryCache{
		data: make(map[string]*entry),
		mu:   &sync.RWMutex{},
	}
	return cache
}

func (c *inMemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.data[key] = &entry{
		value: value,
		timer: time.AfterFunc(ttl, func() {
			c.Delete(key)
		}),
	}
	c.mu.Unlock()
}

func (c *inMemoryCache) Delete(key string) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}

func (c *inMemoryCache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	et, ok := c.data[key]
	if !ok {
		return nil, errors.New("value not found")
	}
	c.mu.RUnlock()
	return et.value, nil
}
