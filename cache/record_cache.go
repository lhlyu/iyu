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

/**
redis design todo
文章浏览总量: hset iyu:article:fire {id} number
            hincrby iyu:article:fire {id} count       // + 1
文章赞总量: hset iyu:article:like {id} number
          hincrby iyu:article:like {id} count       // + 1  or  -1


// 记录用户
文章浏览: set iyu:article:fire:{id} userId
文章Like: set iyu:article:like:{id} userId

*/
