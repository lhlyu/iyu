package po

import "time"

type Record struct {
	Id        int       `json:"id"`                        // 主键
	UserId    int       `json:"userId"`                    // 用户ID
	TargetId  int       `json:"targetId"`                  // 目标ID:文章ID
	Kind      int       `json:"kind"`                      // 记录类型:1-系统;2-异常;3-浏览
	Content   string    `json:"content"`                   // 内容
	Ip        string    `json:"ip" gorm:"default:0.0.0.0"` // IP
	CreatedAt time.Time `json:"createdAt"`                 // 创建时间
}
