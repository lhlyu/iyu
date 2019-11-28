package api

import (
	"fmt"
	"github.com/lhlyu/request"
	"testing"
)

func TestGetArticle(t *testing.T) {
	rq.SetUrl("/article").
		SetParam("id=1").SetMethod(request.GET).
		DoHttp().
		OnSuccess(func(resp request.IResponse) {
			fmt.Println("success:", resp.GetBody())
		}).
		OnError(func(resp request.IResponse) {
			fmt.Println("error:", resp.GetBody())
		})
}
