package cache

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/service/vo"
	"github.com/lhlyu/yutil"
	"strconv"
)

func (c *Cache) SetUser(items []*vo.UserVo) {
	key := c.getKey(userKey)
	c.mutexHandler(key, func() {
		m := make(common.MSF)
		for _, v := range items {
			m[strconv.Itoa(v.Id)] = yutil.JsonObjToStr(v)
		}
		statusCmd := common.Redis.HMSet(key, m)
		if c.Error(statusCmd.Err(), key) {
			return
		}
		common.Redis.Expire(key, day)
	})
}

func (c *Cache) GetUser(fields ...int) []*vo.UserVo {
	key := c.getKey(userKey)
	sliceCmd := common.Redis.HMGet(key, yutil.SliceIntToStr(fields)...)
	if c.Error(sliceCmd.Err(), key, fields) {
		return nil
	}
	vals := sliceCmd.Val()
	var items []*vo.UserVo
	for _, v := range vals {
		if v == nil {
			continue
		}
		item := &vo.UserVo{}
		yutil.JsonStrToObj(v.(string), item)
		items = append(items, item)
	}
	return items
}
