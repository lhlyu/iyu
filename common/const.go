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

// action kind
const (
	KIND_GLOBAL = iota
	KIND_ARTICLE
	KIND_CMNT
	KIND_POST
)

// action
const (
	_ = iota
	ACTION_VISIT
	ACTION_CMNT
	ACTION_LIKE
	ACTION_UNLIKE
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
	ITV   = 3600 * 24
)

// quanta key
const (
	KEY_1 = "admin.pass"
	KEY_2 = "cmnt.open"
	KEY_3 = "cmnt.check"
)
