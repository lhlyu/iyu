package common

const (
	ZERO = iota
	ONE
	TWO
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

const ADMIN = "admin"
