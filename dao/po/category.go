package po

import "time"

type Category struct {
	Id          int       `json:"id"`                     // 主键
	Name        string    `json:"name"`                   // 名字
	Description string    `json:"description"`            // 描述
	Cover       string    `json:"cover"`                  // 封面
	State       int       `json:"state" gorm:"default:1"` // 状态:1-正常;2-已删除
	Number      int       `json:"number"`                 // 关联文章数量
	CreatedAt   time.Time `json:"createdAt"`              // 创建时间
	UpdatedAt   time.Time `json:"updatedAt"`              // 更新时间
}
