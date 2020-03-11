package po

import (
	"github.com/lhlyu/iyu/common"
	"time"
)

// 分类
type Category struct {
	Id          uint      // 主键
	Name        string    // 名字
	Description string    // 描述
	Cover       string    // 封面
	State       uint      // 状态:1-正常;2-已删除
	Number      uint      // 关联文章数量
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间
}

// 获取
func (this *Category) Get() error {
	if this.Id == 0 {
		return MissPkErr
	}
	return common.DB.First(this, this.Id).Error
}

// 分页查询
func (this *Category) Query(rs interface{}, page *common.Page, whr map[string]interface{}, order string) error {
	var total int
	if err := common.DB.Model(this).Where(whr).Count(&total).Error; err != nil {
		return err
	}
	page.SetTotal(total)
	return common.DB.Where(whr).Offset(page.StartRow).Limit(page.PageSize).Order(order).Find(rs).Error
}

// 添加
func (this *Category) Add() error {
	return common.DB.Create(this).Error
}

// 删除
func (this *Category) Del() error {
	if this.Id == 0 {
		return MissPkErr
	}
	return common.DB.Unscoped().Delete(this).Error
}

// 更新
func (this *Category) Update(whr map[string]interface{}) error {
	if this.Id == 0 {
		return MissPkErr
	}
	if whr == nil {
		return common.DB.Model(this).Updates(this).Error
	}
	return common.DB.Model(this).Updates(whr).Error
}
