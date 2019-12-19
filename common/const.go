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
// 1.系统操作;2.错误日志;3.用户登录;4.全站浏览;5.文章浏览;6.文章评论;7.文章回复;8.文章赞;9.文章踩
const (
	_ = iota
	business_system_op
	business_error_log
	business_user_login
	business_global_visit
	business_article_visit
	business_article_comment
	business_article_reply
	business_article_like
	business_article_unlike
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
