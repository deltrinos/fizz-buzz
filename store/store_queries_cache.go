package store

import (
	"github.com/deltrinos/fizz-buzz/models"
	"sync"
)

type CacheResult struct {
	List []string
}

type Cache struct {
	results *sync.Map
}

// Provide caching in the API to allow for faster response times
// based on sync.Map
func (c *Cache) Get(params models.FizzParams) ([]string, bool) {
	res, exists := c.results.Load(params)
	if exists {
		cache, cacheExists := res.(*CacheResult)
		if cacheExists {
			return cache.List, true
		}
	}
	return nil, false
}

func (c *Cache) Save(params models.FizzParams, res []string) []string {
	c.results.Store(params, &CacheResult{
		List: res,
	})
	return res
}
