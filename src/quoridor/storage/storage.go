package storage

import (
	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func Init() {
	c = cache.New(cache.NoExpiration, cache.NoExpiration)
}

func Set(id string, value interface{}) {
	c.Set(id, value, cache.NoExpiration)
}
