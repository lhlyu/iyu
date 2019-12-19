package common

const (
	ZERO = iota
	ONE
	TWO
)

const (
	_         = iota
	UNDELETED // 未删除
	DELETED   // 已删除
)

// business_kind
/**
1.系统操作;2.错误日志;3.用户登录;4.全站浏览;5.文章浏览;6.文章赞;7.文章踩;8.文章评论;9.评论赞;10.评论踩;11.评论回复;12.回复赞;13.回复踩
*/
const (
	_ = iota
	business_system_op
	business_error_log
	business_user_login
	business_global_visit
	business_article_visit
	business_article_like
	business_article_unlike

	business_article_comment
	business_comment_like
	business_comment_unlike

	business_comment_reply
	business_reply_like
	business_reply_unlike
)

// article kind
const (
	ARTICLE_NORMAL = iota + 1
	ARTICLE_ABOUT
	ARTICLE_NOTE
)

const (
	ADMIN = "admin"
	COLOR = "#000000"
	ITV   = 3600 * 24 // 有效时间
)

// quanta key
const (
	KEY_1 = "admin.pass"
	KEY_2 = "cmnt.open"
	KEY_3 = "cmnt.check"
)
