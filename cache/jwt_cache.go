package cache

import (
	"github.com/lhlyu/iyu/common"
	"time"
)

func (c *cache) SetJwt(token string) {
	if c.hasRedis() {
		itv := common.Cfg.GetInt("jwt.itv")
		if itv == 0 {
			itv = common.ITV
		}
		key := common.Cfg.GetString("redis_key.token") + token
		now := time.Now().Format("20060102150405")
		common.Redis.Set(key, now, time.Second*time.Duration(itv))
	}
}

func (c *cache) ExistsJwt(token string) bool {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.token") + token
		val := common.Redis.Exists(key).Val()
		if val > 0 {
			return true
		}
		return false
	}
	return true
}

func (c *cache) DelJwt(token string) {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.token") + token
		common.Redis.Del(key)
	}
}
