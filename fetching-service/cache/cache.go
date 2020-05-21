package cache

import (
	"time"

	"github.com/allegro/bigcache"
)

func NewCacheDefault() (*bigcache.BigCache, error) {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		return nil, err
	}
	return cache, nil

}