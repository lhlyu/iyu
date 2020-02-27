package po

import "time"

type Quanta struct {
	Id          int       `json:"id"`                     // 主键
	Key         string    `json:"key"`                    // key值
	Value       string    `json:"value"`                  // value值
	Description string    `json:"description"`            // 描述
	State       int       `json:"state" gorm:"default:1"` // 状态:1-使用;2-废弃
	CreatedAt   time.Time `json:"createdAt"`              // 创建时间
}
