package cache

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/iyu/util"
	"strconv"
)

const tagKeyName = "redis_key.tag"
const nailKeyName = "redis_key.nail"
const categoryKeyName = "redis_key.category"
const quantaKeyName = "redis_key.quanta"

func (c cache) GetTagAll() []*bo.Tag {
	m := c.GetListData(tagKeyName)
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

func (c cache) GetTag(fields ...int) []*bo.Tag {
	m := c.GetMapData(tagKeyName, util.IntSlinceToStringSlince(fields)...)
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

func (c cache) LoadTags(tags ...*bo.Tag) {
	vm := make(map[string]interface{})
	for _, v := range tags {
		vm[strconv.Itoa(v.Id)] = v
	}
	c.LoadMapAndList(tagKeyName, vm)
}

func (c cache) GetCategoryAll() []*bo.Category {
	m := c.GetListData(categoryKeyName)
	if len(m) == 0 {
		return nil
	}
	var arr []*bo.Category
	for _, v := range m {
		a := &bo.Category{}
		if util.JsonStrToObj(v, a) != nil {
			continue
		}
		arr = append(arr, a)
	}
	return arr
}

func (c cache) GetCategory(fields ...int) []*bo.Category {
	m := c.GetMapData(categoryKeyName, util.IntSlinceToStringSlince(fields)...)
	if len(m) == 0 {
		return nil
	}
	var arr []*bo.Category
	for _, v := range m {
		a := &bo.Category{}
		if util.JsonStrToObj(v, a) != nil {
			continue
		}
		arr = append(arr, a)
	}
	return arr
}

func (c cache) LoadCategorys(categories ...*bo.Category) {
	vm := make(map[string]interface{})
	for _, v := range categories {
		vm[strconv.Itoa(v.Id)] = v
	}
	c.LoadMapAndList(categoryKeyName, vm)
}

func (c cache) GetNailAll() []*bo.Nail {
	m := c.GetListData(nailKeyName)
	if len(m) == 0 {
		return nil
	}
	var arr []*bo.Nail
	for _, v := range m {
		a := &bo.Nail{}
		if util.JsonStrToObj(v, a) != nil {
			continue
		}
		arr = append(arr, a)
	}
	return arr
}

func (c cache) GetNail(fields ...int) []*bo.Nail {
	m := c.GetMapData(nailKeyName, util.IntSlinceToStringSlince(fields)...)
	if len(m) == 0 {
		return nil
	}
	var arr []*bo.Nail
	for _, v := range m {
		a := &bo.Nail{}
		if util.JsonStrToObj(v, a) != nil {
			continue
		}
		arr = append(arr, a)
	}
	return arr
}

func (c cache) LoadNails(nails ...*bo.Nail) {
	vm := make(map[string]interface{})
	for _, v := range nails {
		vm[strconv.Itoa(v.Id)] = v
	}
	c.LoadMapAndList(nailKeyName, vm)
}

func (c cache) GetQuantaPage(page *common.Page) []*bo.Quanta {
	m := c.GetListDataPage(quantaKeyName, page)
	if len(m) == 0 {
		return nil
	}
	var arr []*bo.Quanta
	for _, v := range m {
		a := &bo.Quanta{}
		if util.JsonStrToObj(v, a) != nil {
			continue
		}
		arr = append(arr, a)
	}
	return arr
}

func (c cache) LoadQuantas(quantas ...*bo.Quanta) {
	vm := make(map[string]interface{})
	for _, v := range quantas {
		vm[v.Key] = v
	}
	c.LoadMapAndList(quantaKeyName, vm)
}
