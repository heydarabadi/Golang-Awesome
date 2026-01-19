package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]string
	lock sync.RWMutex
}

func (c *Cache) Get(key string) (string, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	val, ok := c.data[key]
	return val, ok
}
func (c *Cache) Set(key, value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = value
}
func slowFunction(key string) string {
	time.Sleep(2 * time.Second)
	return "Result for " + key
}
func main() {
	cache := Cache{data: make(map[string]string)}
	key := "abc"
	// First fetch (cache miss)
	if val, ok := cache.Get(key); ok {
		fmt.Println("From cache:", val)
	} else {
		fmt.Println("Cache miss. Calling slowFunction...")
		val := slowFunction(key)
		cache.Set(key, val)
		fmt.Println("New cache value:", val)
	}
	// Second fetch (cache hit)
	if val, ok := cache.Get(key); ok {
		fmt.Println("From cache:", val)
	}
}
