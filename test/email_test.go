package test

import (
	"github.com/lhlyu/iyu/common"
	"testing"
)

func TestEmail(t *testing.T) {
	common.Email.Send(common.NewMessageContent(
		"yu",
		"sadasdasd23d3e223dd",
	))
}
