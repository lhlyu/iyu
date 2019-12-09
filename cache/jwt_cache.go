package cache

import (
	"github.com/lhlyu/iyu/common"
	"strconv"
	"time"
)

func (c *cache) SetJwt(userId int, token string) {
	if c.hasRedis() {
		itv := common.Cfg.GetInt("jwt.itv")
		if itv == 0 {
			itv = common.ITV
		}
		key := common.Cfg.GetString("redis_key.token") + strconv.Itoa(userId)
		err := common.Redis.Set(key, token, time.Second*time.Duration(itv)).Err()
		c.Error(err)
	}
}

func (c *cache) GetJwt(userId int) string {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.token") + strconv.Itoa(userId)
		return common.Redis.Get(key).Val()
	}
	return ""
}

func (c *cache) DelJwt(userId int) {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.token") + strconv.Itoa(userId)
		err := common.Redis.Del(key).Err()
		c.Error(err)
	}
}
