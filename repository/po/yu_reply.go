package po

import "time"

type YuReply struct {
	Id        int       `json:"id"         db:"id"`
	ArticleId int       `json:"articleId"  db:"article_id"` // 文章ID
	CmntId    int       `json:"cmntId"     db:"cmnt_id"`    // 评论ID
	SortNum   int       `json:"sortNum"    db:"sort_num"`   // 排序
	Floor     int       `json:"floor"      db:"floor"`      // 回复楼层
	Like      int       `json:"like"       db:"like"`       // 赞
	IsDelete  int       `json:"isDelete"   db:"is_delete"`  // 删除:1-未删除;2-已删除
	IsCheck   int       `json:"isCheck"    db:"is_check"`   // 审核:1-未审核;2-已审核
	Content   string    `json:"content"    db:"content"`    // 内容
	AtUserId  int       `json:"atUserId"   db:"at_user_id"` // 艾特用户ID
	AtFloor   int       `json:"atFloor"    db:"at_floor"`   // 艾特楼层
	CreatedAt time.Time `json:"createdAt"  db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"  db:"updated_at"`
}
