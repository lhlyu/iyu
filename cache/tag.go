package cache

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/service/vo"
	"github.com/lhlyu/yutil"
	"strconv"
)

func (c *Cache) SetTag(items []*vo.TagVo) {
	key := c.getKey(tagKey)
	c.mutexHandler(key, func() {
		m := make(common.MSF)
		for _, v := range items {
			m[strconv.Itoa(v.Id)] = yutil.JsonObjToStr(v)
		}
		statusCmd := common.Redis.HMSet(key, m)
		if c.Error(statusCmd.Err(), key) {
			return
		}
		common.Redis.Expire(key, week)
	})
}

func (c *Cache) GetTag(fields ...int) []*vo.TagVo {
	key := c.getKey(tagKey)
	sliceCmd := common.Redis.HMGet(key, yutil.SliceIntToStr(fields)...)
	if c.Error(sliceCmd.Err(), key, fields) {
		return nil
	}
	vals := sliceCmd.Val()
	var items []*vo.TagVo
	for _, v := range vals {
		item := &vo.TagVo{}
		yutil.JsonStrToObj(v.(string), item)
		items = append(items, item)
	}
	return items
}
