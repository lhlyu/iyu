package po

import "time"

type User struct {
	Id        int       `json:"id"`                        // 主键
	Account   string    `json:"account"`                   // 账号
	Name      string    `json:"name"`                      // 昵称
	Uuid      string    `json:"uuid"`                      // 唯一ID
	Source    string    `json:"source"`                    // 来源
	Avatar    string    `json:"avatar"`                    // 头像
	Url       string    `json:"url"`                       // 来源个人主页
	Website   string    `json:"website"`                   // 个人网站
	Bio       string    `json:"bio"`                       // 个性签名
	Bg        string    `json:"bg"`                        // 背景图
	Location  string    `json:"location"`                  // 位置
	Kind      int       `json:"kind" gorm:"default:1"`     // 用户类型:1-普通;2-好友
	State     int       `json:"state" gorm:"default:1"`    // 用户状态:1-正常;2-已删除
	Ip        string    `json:"ip" gorm:"default:0.0.0.0"` // 最近访问Ip
	CreatedAt time.Time `json:"createdAt"`                 // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`                 // 更新时间
}
