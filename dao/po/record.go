package po

import (
	"github.com/lhlyu/iyu/common"
	"time"
)

// 记录表:我来到，我看到，我记录
type Record struct {
	Id        uint      // 主键
	UserId    uint      // 用户ID
	TargetId  uint      // 目标ID:文章ID
	Kind      uint      // 记录类型:1-系统;2-异常;3-浏览
	Content   string    // 内容
	Ip        string    // IP
	CreatedAt time.Time // 创建时间
}

// 获取
func (this *Record) Get() error {
	if this.Id == 0 {
		return MissPkErr
	}
	return common.DB.First(this, this.Id).Error
}

// 分页查询
func (this *Record) Query(rs interface{}, page *common.Page, whr map[string]interface{}, order string) error {
	var total int
	if err := common.DB.Model(this).Where(whr).Count(&total).Error; err != nil {
		return err
	}
	page.SetTotal(total)
	return common.DB.Where(whr).Offset(page.StartRow).Limit(page.PageSize).Order(order).Find(rs).Error
}

// 添加
func (this *Record) Add() error {
	return common.DB.Create(this).Error
}

// 删除
func (this *Record) Del() error {
	if this.Id == 0 {
		return MissPkErr
	}
	return common.DB.Unscoped().Delete(this).Error
}

// 更新
func (this *Record) Update(whr map[string]interface{}) error {
	if this.Id == 0 {
		return MissPkErr
	}
	if whr == nil {
		return common.DB.Model(this).Updates(this).Error
	}
	return common.DB.Model(this).Updates(whr).Error
}
