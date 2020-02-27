package po

import "time"

type Plate struct {
	Id          int       `json:"id"`                     // 主键
	Name        string    `json:"name"`                   // 名字
	Description string    `json:"description"`            // 描述
	State       int       `json:"state" gorm:"default:1"` // 状态:1-开启;2-关闭
	CreatedAt   time.Time `json:"createdAt"`              // 创建时间
	UpdatedAt   time.Time `json:"updatedAt"`              // 更新时间
}
