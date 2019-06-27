package storage

import (
	"strconv"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func new() {
	c = cache.New(-1, -1)
}

func Set(id int, value interface{}) {
	if c == nil {
		new()
	}
	c.Set(strconv.Itoa(id), value, cache.NoExpiration)
}
