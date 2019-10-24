package po

import "time"

type YuUserInfo struct {
    Id         int        `db:"id"`          
    UserId     int        `db:"user_id"`      // 用户ID
    AvatarUrl  string     `db:"avatar_url"`   // 用户头像
    UserUrl    string     `db:"user_url"`     // 用户地址
    UserName   string     `db:"user_name"`    // 用户名字
    Bio        string     `db:"bio"`          // 个性签名
    Ip         string     `db:"ip"`           // ip
    CreatedAt  time.Time  `db:"created_at"`  
    UpdatedAt  time.Time  `db:"updated_at"`  
}

