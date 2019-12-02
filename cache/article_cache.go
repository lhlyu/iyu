package cache

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/iyu/util"
	"strconv"
)

func (c cache) LoadArticle(articleData *bo.ArticleData) {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.article")
		if key == "" {
			return
		}
		key += _MAP
		c.mutexHandler(key, func() {
			common.Redis.HSet(key, strconv.Itoa(articleData.ID), util.ObjToJsonStr(articleData))
		})
	}
}

func (c cache) GetArticles(fields ...int) []*bo.ArticleData {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.article")
		if key == "" {
			return nil
		}
		key += _MAP
		interArr := common.Redis.HMGet(key, util.IntSlinceToStringSlince(fields)...).Val()
		var arr []*bo.ArticleData
		for _, v := range interArr {
			if v == nil {
				continue
			}
			a := &bo.ArticleData{}
			if util.JsonStrToObj(fmt.Sprint(v), a) != nil {
				continue
			}
			arr = append(arr, a)
		}
		return arr
	}
	return nil
}

func (c cache) DelArticles(fields ...int) {
	if c.hasRedis() {
		key := common.Cfg.GetString("redis_key.article")
		if key == "" {
			return
		}
		key += _MAP
		common.Redis.HDel(key, util.IntSlinceToStringSlince(fields)...).Val()
	}
}
