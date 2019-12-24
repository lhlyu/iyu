package cache

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/service/vo"
	"github.com/lhlyu/yutil"
	"strconv"
)

func (c *Cache) SetCategory(items []*vo.CategoryVo) {
	key := c.getKey(categoryKey)
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

func (c *Cache) GetCategory(fields ...int) []*vo.CategoryVo {
	key := c.getKey(categoryKey)
	sliceCmd := common.Redis.HMGet(key, yutil.SliceIntToStr(fields)...)
	if c.Error(sliceCmd.Err(), key, fields) {
		return nil
	}
	vals := sliceCmd.Val()
	var items []*vo.CategoryVo
	for _, v := range vals {
		if v == nil {
			continue
		}
		item := &vo.CategoryVo{}
		yutil.JsonStrToObj(v.(string), item)
		items = append(items, item)
	}
	return items
}
