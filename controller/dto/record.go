package dto

type RecordDto struct {
	PageNum      int    `json:"pageNum" validate:"required,gt=0"`  // 当前页码
	PageSize     int    `json:"pageSize" validate:"required,gt=0"` // 每页记录条数
	UserId       int    `json:"userId"`                            // 用户ID
	BusinessId   int    `json:"businessId"`                        // 目标Id
	Content      string `json:"content"`                           // 内容
	BusinessKind int    `json:"businessKind"`                      // 1.系统操作;2.错误日志;3.用户登录;4.全站浏览;5.文章浏览;6.文章赞;7.文章踩;8.文章评论;9.评论赞;10.评论踩;11.评论回复;12.回复赞;13.回复踩
	Agent        string `json:"agent"`
	Ip           string `json:"ip"`
	CreatedAt    string `json:"createdAt"`
}
