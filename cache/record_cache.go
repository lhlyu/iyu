package cache

import (
	"github.com/lhlyu/iyu/common"
)

func (c *cache) Record(key string, f func()) {
	if c.hasRedis() {
		keyMutex := common.Cfg.GetString("redis_key.iyu") + key + ":mutex"
		if rs, err := common.Redis.SetNX(keyMutex, 1, _ONE_DAY).Result(); !rs {
			c.Error(err)
			return
		}
		f()
	}
}
