package cache

import (
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/yutil"
	"strconv"
)

const (
	_tag_key      = "redis_key.tag"
	_category_key = "redis_key.category"
	_nail_key     = "redis_key.nail"
	_quanta_map   = "redis_key.quanta"
	_quanta_key   = "redis_key.quanta_key"
	_user_key     = "redis_key.user"
	_article_key  = "redis_key.article"
)

// tag
func (c cache) SetTag(values ...*bo.Tag) {
	if len(values) == 0 {
		return
	}
	if c.hasRedis() {
		m := make(map[string]interface{})
		for _, v := range values {
			value := yutil.JsonObjToStr(v)
			m[strconv.Itoa(v.Id)] = value
		}
		c.setMap(_tag_key, m, _ONE_WEEK)
	}
}

func (c cache) GetTag(field ...int) []*bo.Tag {
	if c.hasRedis() {
		arr := c.getMap(_tag_key, yutil.SliceIntToStr(field)...)
		var values []*bo.Tag
		for _, v := range arr {
			value := &bo.Tag{}
			yutil.JsonStrToObj(v, value)
			if value == nil {
				continue
			}
			values = append(values, value)
		}
		return values
	}
	return nil
}

// category
func (c cache) SetCategory(values ...*bo.Category) {
	if len(values) == 0 {
		return
	}
	if c.hasRedis() {
		m := make(map[string]interface{})
		for _, v := range values {
			value := yutil.JsonObjToStr(v)
			m[strconv.Itoa(v.Id)] = value
		}
		c.setMap(_category_key, m, _ONE_WEEK)
	}
}

func (c cache) GetCategory(field ...int) []*bo.Category {
	if c.hasRedis() {
		arr := c.getMap(_category_key, yutil.SliceIntToStr(field)...)
		var values []*bo.Category
		for _, v := range arr {
			value := &bo.Category{}
			yutil.JsonStrToObj(v, value)
			if value == nil {
				continue
			}
			values = append(values, value)
		}
		return values
	}
	return nil
}

// nail
func (c cache) SetNail(values ...*bo.Nail) {
	if len(values) == 0 {
		return
	}
	if c.hasRedis() {
		m := make(map[string]interface{})
		for _, v := range values {
			value := yutil.JsonObjToStr(v)
			m[strconv.Itoa(v.Id)] = value
		}
		c.setMap(_nail_key, m, _ONE_WEEK)
	}
}

func (c cache) GetNail(field ...int) []*bo.Nail {
	if c.hasRedis() {
		arr := c.getMap(_nail_key, yutil.SliceIntToStr(field)...)
		var values []*bo.Nail
		for _, v := range arr {
			value := &bo.Nail{}
			yutil.JsonStrToObj(v, value)
			if value == nil {
				continue
			}
			values = append(values, value)
		}
		return values
	}
	return nil
}

// quanta
func (c cache) SetQuanta(values ...*bo.Quanta) {
	if len(values) == 0 {
		return
	}
	if c.hasRedis() {
		m := make(map[string]interface{})
		n := make(map[string]interface{})
		for _, v := range values {
			value := yutil.JsonObjToStr(v)
			m[strconv.Itoa(v.Id)] = value
			n[v.Key] = value
		}
		c.setMap(_quanta_map, m, _ONE_WEEK)
		c.setMap(_quanta_key, n, _ONE_WEEK)
	}
}

func (c cache) GetQuanta(field ...int) []*bo.Quanta {
	if c.hasRedis() {
		arr := c.getMap(_quanta_map, yutil.SliceIntToStr(field)...)
		var values []*bo.Quanta
		for _, v := range arr {
			value := &bo.Quanta{}
			yutil.JsonStrToObj(v, value)
			if value == nil {
				continue
			}
			values = append(values, value)
		}
		return values
	}
	return nil
}

func (c cache) GetQuantaByKeys(field ...string) []*bo.Quanta {
	if c.hasRedis() {
		arr := c.getMap(_quanta_key, field...)
		var values []*bo.Quanta
		for _, v := range arr {
			value := &bo.Quanta{}
			yutil.JsonStrToObj(v, value)
			if value == nil {
				continue
			}
			values = append(values, value)
		}
		return values
	}
	return nil
}

// user
func (c cache) SetUser(values ...*bo.User) {
	if len(values) == 0 {
		return
	}
	if c.hasRedis() {
		m := make(map[string]interface{})
		for _, v := range values {
			value := yutil.JsonObjToStr(v)
			m[strconv.Itoa(v.Id)] = value
		}
		c.setMap(_user_key, m, _ONE_WEEK)
	}
}

func (c cache) GetUser(field ...int) []*bo.User {
	if c.hasRedis() {
		arr := c.getMap(_user_key, yutil.SliceIntToStr(field)...)
		var values []*bo.User
		for _, v := range arr {
			value := &bo.User{}
			yutil.JsonStrToObj(v, value)
			if value == nil {
				continue
			}
			values = append(values, value)
		}
		return values
	}
	return nil
}

// article
func (c cache) SetArticle(values ...*bo.Article) {
	if len(values) == 0 {
		return
	}
	if c.hasRedis() {
		m := make(map[string]interface{})
		for _, v := range values {
			value := yutil.JsonObjToStr(v)
			m[strconv.Itoa(v.Id)] = value
		}
		c.setMap(_article_key, m, _ONE_WEEK)
	}
}

func (c cache) GetArticle(field ...int) []*bo.Article {
	if c.hasRedis() {
		arr := c.getMap(_article_key, yutil.SliceIntToStr(field)...)
		var values []*bo.Article
		for _, v := range arr {
			value := &bo.Article{}
			yutil.JsonStrToObj(v, value)
			if value == nil {
				continue
			}
			values = append(values, value)
		}
		return values
	}
	return nil
}
