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
		listKey := key + _LIST
		targetListKey := common.Redis.Get(listKey).Val()
		if targetListKey == "" {
			return nil
		}
		m := common.Redis.LRange(targetListKey, 0, -1).Val()
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

func (c cache) LoadTags(tags ...*bo.Tag) {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.tag")
		if key == "" {
			return
		}
		mapKey := key + _MAP
		targetMapKey := mapKey + c.getTimestamp()
		listKey := key + _LIST
		targetListKey := listKey + c.getTimestamp()
		c.mutexHandler(mapKey, func() {
			var arr []interface{}
			for _, v := range tags {
				value := util.ObjToJsonStr(v)
				common.Redis.HSet(targetMapKey, strconv.Itoa(v.Id), value)
				arr = append(arr, value)
			}
			common.Redis.Expire(targetMapKey, _ONE_WEEK)
			oldTargetMapKey := common.Redis.Get(mapKey).Val()
			common.Redis.Set(mapKey, targetMapKey, _ONE_WEEK)
			if oldTargetMapKey != "" {
				common.Redis.Del(oldTargetMapKey)
			}
			if len(arr) > 0 {
				common.Redis.RPush(targetListKey, arr...)
			}
			oldTargetListKey := common.Redis.Get(listKey).Val()
			common.Redis.Set(listKey, targetListKey, _ONE_WEEK)
			if oldTargetListKey != "" {
				common.Redis.Del(oldTargetListKey)
			}
		})

	}
}
