package model

type YuRecord struct {
    Id         int        `db:"id"`          
    UserId     int        `db:"user_id"`     
    ArticleId  int        `db:"article_id"`  
    Action     int        `db:"action"`       // 动作:1-浏览;2-评论
    Ip         string     `db:"ip"`          
    CreatedAt  int        `db:"created_at"`
    UpdatedAt  int        `db:"updated_at"`
}

