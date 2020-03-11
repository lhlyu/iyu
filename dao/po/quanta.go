package po

import (
	"github.com/lhlyu/iyu/common"
	"time"
)

// 配置表
type Quanta struct {
	Id          uint      // 主键
	Key         string    // key值
	Value       string    // value值
	Description string    // 描述
	State       uint      // 状态:1-使用;2-废弃
	CreatedAt   time.Time // 创建时间
}

// 获取
func (this *Quanta) Get() error {
	if this.Id == 0 {
		return MissPkErr
	}
	return common.DB.First(this, this.Id).Error
}

// 分页查询
func (this *Quanta) Query(rs interface{}, page *common.Page, whr map[string]interface{}, order string) error {
	var total int
	if err := common.DB.Model(this).Where(whr).Count(&total).Error; err != nil {
		return err
	}
	page.SetTotal(total)
	return common.DB.Where(whr).Offset(page.StartRow).Limit(page.PageSize).Order(order).Find(rs).Error
}

// 添加
func (this *Quanta) Add() error {
	return common.DB.Create(this).Error
}

// 删除
func (this *Quanta) Del() error {
	if this.Id == 0 {
		return MissPkErr
	}
	return common.DB.Unscoped().Delete(this).Error
}

// 更新
func (this *Quanta) Update(whr map[string]interface{}) error {
	if this.Id == 0 {
		return MissPkErr
	}
	if whr == nil {
		return common.DB.Model(this).Updates(this).Error
	}
	return common.DB.Model(this).Updates(whr).Error
}
