package cache

import (
	"errors"
	"sync"
)

type Cache interface {
	Set(key string, value interface{})
	Delete(key string)
	Get(key string) (interface{}, error)
}

type inMemoryCache struct {
	data map[string]interface{}
	mu   *sync.RWMutex
}

func New() Cache {
	return &inMemoryCache{
		data: make(map[string]interface{}),
		mu:   &sync.RWMutex{},
	}
}

func (c *inMemoryCache) Set(key string, value interface{}) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

func (c *inMemoryCache) Delete(key string) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}

func (c *inMemoryCache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	if !ok {
		return nil, errors.New("value not found")
	}
	return value, nil
}
