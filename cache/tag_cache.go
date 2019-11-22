package cache

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/iyu/util"
	"strconv"
)

func (c cache) GetTagAll() []*bo.Tag {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.tag")
		if key == "" {
			return nil
		}
		m := common.Redis.HGetAll(key).Val()
		if len(m) == 0 {
			return nil
		}
		var arr []*bo.Tag
		for _, v := range m {
			a := &bo.Tag{}
			if util.JsonStrToObj(v, a) != nil {
				continue
			}
			arr = append(arr, a)
		}
		return arr
	}
	return nil
}

func (c cache) GetTags(ids ...int) []*bo.Tag {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.tag")
		if key == "" {
			return nil
		}
		idStr := util.IntSlinceToStringSlince(ids)
		m := common.Redis.HMGet(key, idStr...).Val()
		if len(m) == 0 {
			return nil
		}
		var arr []*bo.Tag
		for _, v := range m {
			a := v.(*bo.Tag)
			arr = append(arr, a)
		}
		return arr
	}
	return nil
}

func (c cache) LoadTags(tags []*bo.Tag) {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.tag")
		if key == "" {
			return
		}
		for _, v := range tags {
			common.Redis.HSet(key, strconv.Itoa(v.Id), util.ObjToJsonStr(v))
		}
		common.Redis.Expire(key, _ONE_MONTH)
	}

}
