package cache

import (
	"fmt"
	"sync"
)

type Cache struct {
	sync.RWMutex
	RequestsPerUser map[string]int
	Bots            int
}

var CacheApp *Cache

func New() *Cache {
	if CacheApp == nil {
		CacheApp = &Cache{
			RequestsPerUser: make(map[string]int),
			Bots:            0,
		}
	}
	return CacheApp
}

func (c *Cache) Set(userId string) {
	c.Lock()
	defer c.Unlock()
	c.RequestsPerUser[userId]++
}

func (c *Cache) Get(key string) int {
	c.RLock()
	defer c.RUnlock()
	if val, ok := c.RequestsPerUser[key]; ok {
		return val
	} else {
		//perhaps need to cover id more detail
		fmt.Println(" Key Not Present ")
		return 0
	}
}

func (c *Cache) GetBotsCount() int {
	c.RLock()
	defer c.RUnlock()
	return c.Bots
}
