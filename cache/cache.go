package cache

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/trace"
	"github.com/lhlyu/yutil/v2"
	"time"
)

type BaseCache struct {
	Prefix string
	trace.ITracker
}

func NewBaseCache(tracker trace.ITracker) BaseCache {
	prefix := common.Cfg.GetString("redis_key.prefix")
	return BaseCache{prefix, tracker}
}

func (c BaseCache) IsExists(key string) bool {
	initCmd := common.Redis.Exists(key)
	if initCmd.Err() != nil {
		c.Error(initCmd.Err(), key)
		return false
	}
	return initCmd.Val() > 0
}

func (c BaseCache) Key(key string) string {
	val := common.Cfg.GetString(key)
	if val == "" {
		return c.Prefix + key
	}
	return c.Prefix + val
}

func (c BaseCache) Lock(name string, fn func()) {
	if name == "" {
		name = c.Prefix
	}
	key := name + ":lock"
	ok := false
	for i := 0; i < 3; i++ {
		boolCmd := common.Redis.SetNX(key, "1", time.Second*10)
		if err := boolCmd.Err(); err != nil {
			c.Error(err, key)
			return
		}
		if ok = boolCmd.Val(); ok {
			break
		}
		time.Sleep(time.Second)
	}
	if !ok {
		return
	}
	fn()
	common.Redis.Expire(key, 0)
}

func (c *BaseCache) Set(key string, v interface{}, expiration time.Duration) {
	c.Lock(key, func() {
		common.Redis.Set(key, yutil.Json.Marshal(v), expiration)
	})
}

func (c *BaseCache) Get(key string, dist interface{}) (bool, error) {
	if c.IsExists(key) {
		stringCmd := common.Redis.Get(key)
		if err := stringCmd.Err(); err != nil {
			c.Error(err, key)
			return false, err
		}
		yutil.Json.Unmarshal(stringCmd.Val(), dist)
		return true, nil
	}
	return false, nil
}

func (c *BaseCache) Clear(key string) {
	c.Lock(key, func() {
		if c.IsExists(key) {
			boolCmd := common.Redis.Expire(key, 0)
			if boolCmd.Err() != nil {
				c.Error(boolCmd.Err(), key)
			}
		}
	})
}
