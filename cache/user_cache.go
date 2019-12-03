package cache

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/iyu/util"
	"strconv"
)

func (c cache) LoadUser(articleDatas []*bo.UserData) {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.user")
		if key == "" {
			return
		}
		key += _MAP
		c.mutexHandler(key, func() {
			for _, v := range articleDatas {
				common.Redis.HSet(key, strconv.Itoa(v.Id), util.ObjToJsonStr(v))
			}
			common.Redis.Expire(key, _ONE_MONTH)
		})
	}
}

func (c cache) GetUsers(fields ...int) []*bo.UserData {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.user")
		if key == "" {
			return nil
		}
		key += _MAP
		interArr := common.Redis.HMGet(key, util.IntSlinceToStringSlince(fields)...).Val()
		var arr []*bo.UserData
		for _, v := range interArr {
			if v == nil {
				continue
			}
			a := &bo.UserData{}
			if util.JsonStrToObj(fmt.Sprint(v), a) != nil {
				continue
			}
			arr = append(arr, a)
		}
		return arr
	}
	return nil
}

func (c cache) DelUsers(fields ...int) {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.user")
		if key == "" {
			return
		}
		key += _MAP
		common.Redis.HDel(key, util.IntSlinceToStringSlince(fields)...).Val()
	}
}
