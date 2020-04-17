package po

import (
	"github.com/lhlyu/iyu/common"
	"time"
)

type Article struct {
	Id           uint      `db:"id"`            // 自增主键
	Code         string    `db:"code"`          // 唯一码
	Title        string    `db:"title"`         // 标题
	Toc          string    `db:"toc"`           // 目录
	Summary      string    `db:"summary"`       // 摘要
	Content      string    `db:"content"`       // 内容
	Cover        string    `db:"cover"`         // 封面
	Labels       string    `db:"labels"`        // 标签
	Kind         string    `db:"kind"`          // 类型:normal(普通),gist(灵感),self(自述)
	State        string    `db:"state"`         // 状态:draft(草稿),private(私密发布),publish(开放发布),dustbin(垃圾箱)
	Password     string    `db:"password"`      // 密码,配合私密发布
	CommentState string    `db:"comment_state"` // 评论状态:open(开放),close(关闭),owner(仅所有者)
	Remake       string    `db:"remake"`        // 备注
	Sort         uint      `db:"sort"`          // 排序,降序
	Category     uint      `db:"category"`      // 分类
	CommentCount uint      `db:"comment_count"` // 评论数量
	ViewCount    uint      `db:"view_count"`    // 访问数量
	GoodCount    uint      `db:"good_count"`    // 点赞数量
	BadCount     uint      `db:"bad_count"`     // 踩数量
	CreatedAt    time.Time `db:"created_at"`    // 创建时间
	UpdatedAt    time.Time `db:"updated_at"`    // 更新时间
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
