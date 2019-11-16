package model

type YuArticle struct {
	Id            int    `db:"id"`             // 文章ID
	UserId        int    `db:"user_id"`        // 用户ID
	Like          int    `db:"like"`           // 赞
	Unlike        int    `db:"unlike"`         // 踩
	Wraper        string `db:"wraper"`         // 头背景
	Title         string `db:"title"`          // 标题
	Content       string `db:"content"`        // 内容
	IsTop         int    `db:"is_top"`         // 是否置顶:1-不置顶;2-置顶
	CategoryId    int    `db:"category_id"`    // 分类ID
	NailId        int    `db:"nail_id"`        // 钉子ID
	Kind          int    `db:"kind"`           // 文章类型:1-普通文章;2-特殊文章
	IsDelete      int    `db:"is_delete"`      // 是否已删除:1-未删除;2-已删除
	CreatedAt     int    `db:"created_at"`     // 创建时间
	UpdatedAt     int    `db:"updated_at"`     // 更新时间
}
