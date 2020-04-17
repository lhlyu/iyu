package cache

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/trace"
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
