package po

import (
	"github.com/lhlyu/iyu/common"
	"time"
)

// 用户表:光孕众生，众生随影
type User struct {
	Id        uint      // 主键
	Account   string    // 账号
	Name      string    // 昵称
	Uuid      string    // 唯一ID
	Source    string    // 来源
	Avatar    string    // 头像
	Url       string    // 来源个人主页
	Website   string    // 个人网站
	Bio       string    // 个性签名
	Bg        string    // 背景图
	Location  string    // 位置
	Kind      uint      // 用户类型:1-普通;2-好友
	State     uint      // 用户状态:1-正常;2-已删除
	Ip        string    // 最近访问Ip
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}

// 获取
func (this *User) Get() error {
	if this.Id == 0 {
		return MissPkErr
	}
	return common.DB.First(this, this.Id).Error
}

// 分页查询
func (this *User) Query(rs interface{}, page *common.Page, whr map[string]interface{}, order string) error {
	var total int
	if err := common.DB.Model(this).Where(whr).Count(&total).Error; err != nil {
		return err
	}
	page.SetTotal(total)
	return common.DB.Where(whr).Offset(page.StartRow).Limit(page.PageSize).Order(order).Find(rs).Error
}

// 添加
func (this *User) Add() error {
	return common.DB.Create(this).Error
}

// 删除
func (this *User) Del() error {
	if this.Id == 0 {
		return MissPkErr
	}
	return common.DB.Unscoped().Delete(this).Error
}

// 更新
func (this *User) Update(whr map[string]interface{}) error {
	if this.Id == 0 {
		return MissPkErr
	}
	if whr == nil {
		return common.DB.Model(this).Updates(this).Error
	}
	return common.DB.Model(this).Updates(whr).Error
}
