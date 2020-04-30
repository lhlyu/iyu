package cache

import (
	"github.com/lhlyu/iyu/trace"
	"time"
)

type CategoryCache struct {
	BaseCache
}

func NewCategoryCache(tracker trace.ITracker) CategoryCache {
	return CategoryCache{
		BaseCache: NewBaseCache(tracker),
	}
}

func (c CategoryCache) Set(v interface{}) {
	key := c.Key("redis_key.category")
	c.BaseCache.Set(key, v, time.Second*600)
}

func (c CategoryCache) Get(dist interface{}) (bool, error) {
	key := c.Key("redis_key.category")
	return c.BaseCache.Get(key, dist)
}

func (c CategoryCache) Clear() {
	key := c.Key("redis_key.category")
	c.BaseCache.Clear(key)
}
