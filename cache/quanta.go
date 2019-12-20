package cache

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/service/vo"
	"github.com/lhlyu/yutil"
)

func (c *Cache) SetQuanta(items []*vo.QuantaVo) {
	key := c.getKey(quantaKey)
	c.mutexHandler(key, func() {
		m := make(common.MSF)
		for _, v := range items {
			m[v.Key] = yutil.JsonObjToStr(v)
		}
		statusCmd := common.Redis.HMSet(key, m)
		if c.Error(statusCmd.Err(), key) {
			return
		}
		common.Redis.Expire(key, week)
	})
}

func (c *Cache) GetQuanta(fields ...string) []*vo.QuantaVo {
	key := c.getKey(quantaKey)
	sliceCmd := common.Redis.HMGet(key, fields...)
	if c.Error(sliceCmd.Err(), key, fields) {
		return nil
	}
	vals := sliceCmd.Val()
	var items []*vo.QuantaVo
	for _, v := range vals {
		item := &vo.QuantaVo{}
		yutil.JsonStrToObj(v.(string), item)
		items = append(items, item)
	}
	return items
}
