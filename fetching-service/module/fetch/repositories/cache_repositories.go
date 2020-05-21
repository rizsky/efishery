package repositories

import (
	"fmt"
	"strconv"

	"github.com/allegro/bigcache"
)

//CacheRepo :struct
type CacheRepo struct {
	cache *bigcache.BigCache
}

//NewCacheRepo init function
func NewCacheRepo(cache *bigcache.BigCache) CacheRepository {
	return &CacheRepo{
		cache: cache,
	}
}

//Set usd idr for
func (c *CacheRepo) Set(key string, value float64) error {
	v := fmt.Sprintf("%f", value)

	err := c.cache.Set(key, []byte(v))
	if err != nil {
		return err
	}
	return nil

}

//Get usd idr
func (c *CacheRepo) Get(key string) (float64, error) {
	entry, err := c.cache.Get(key)
	if err == bigcache.ErrEntryNotFound {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(string(entry), 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}
