package po

import (
	"github.com/lhlyu/iyu/common"
	"time"
)

// 文章表
type Article struct {
	Id         uint      // 主键
	Code       string    // 文章code
	PlateId    uint      // 所属板块ID
	IsTop      uint      // 是否置顶:1-否;2-是
	CategoryId uint      // 分类ID
	Color      string    // 颜色
	Labels     string    // 标签
	Title      string    // 标题
	Summary    string    // 摘要
	Content    string    // 内容
	Cover      string    // 图片
	State      uint      // 状态:1-正常;2-关闭
	CreatedAt  time.Time // 创建时间
	UpdatedAt  time.Time // 修改时间
}

// 获取
func (this *Article) Get() error {
	if this.Id == 0 {
		return MissPkErr
	}
	return common.DB.First(this, this.Id).Error
}

// 分页查询
func (this *Article) Query(rs interface{}, page *common.Page, whr map[string]interface{}, order string) error {
	var total int
	if err := common.DB.Model(this).Where(whr).Count(&total).Error; err != nil {
		return err
	}
	page.SetTotal(total)
	return common.DB.Where(whr).Offset(page.StartRow).Limit(page.PageSize).Order(order).Find(rs).Error
}

// 添加
func (this *Article) Add() error {
	return common.DB.Create(this).Error
}

// 删除
func (this *Article) Del() error {
	if this.Id == 0 {
		return MissPkErr
	}
	return common.DB.Unscoped().Delete(this).Error
}

// 更新
func (this *Article) Update(whr map[string]interface{}) error {
	if this.Id == 0 {
		return MissPkErr
	}
	if whr == nil {
		return common.DB.Model(this).Updates(this).Error
	}
	return common.DB.Model(this).Updates(whr).Error
}
