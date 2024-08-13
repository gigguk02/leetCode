package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[k] = v
}

func (c *Cache) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[k]
	return v, ok
}
func main() {
	cache := NewCache()

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	if value, ok := cache.Get("key1"); ok {
		fmt.Println("Value for key1:", value)
	} else {
		fmt.Println("Key1 not found")
	}

	if value, ok := cache.Get("key2"); ok {
		fmt.Println("Value for key2:", value)
	} else {
		fmt.Println("Key2 not found")
	}

	if _, ok := cache.Get("key3"); !ok {
		fmt.Println("Key3 not found")
	}
}
