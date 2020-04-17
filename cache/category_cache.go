package cache

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/trace"
	"github.com/lhlyu/yutil/v2"
	"time"
)

type CategoryCache struct {
	BaseCache
}

func NewCategoryCache(tracker trace.ITracker) *CategoryCache {
	return &CategoryCache{
		BaseCache: NewBaseCache(tracker),
	}
}

func (c *CategoryCache) Set(v interface{}) {
	key := c.Key("redis_key.category")
	common.Redis.Set(key, yutil.Json.Marshal(v), time.Second*60)
}

func (c *CategoryCache) Get(dist interface{}) (bool, error) {
	key := c.Key("redis_key.category")
	if c.IsExists(key) {
		stringCmd := common.Redis.Get(key)
		if err := stringCmd.Err(); err != nil {
			c.Error(err, key)
			return false, err
		}
		yutil.Json.Unmarshal(stringCmd.String(), dist)
		return true, nil
	}
	return false, nil
}

func (c *CategoryCache) Clear() {
	key := c.Key("redis_key.category")
	if c.IsExists(key) {
		initCmd := common.Redis.Del(key)
		if initCmd.Err() != nil {
			c.Error(initCmd.Err(), key)
		}
	}
}
