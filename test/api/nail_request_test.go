package api

import (
	"fmt"
	"github.com/lhlyu/request"
	"testing"
)

func TestGetAllNail(t *testing.T) {
	rq.SetUrl("/category").
		SetMethod(request.GET).
		DoHttp().
		OnSuccess(func(resp request.IResponse) {
			fmt.Println("success:", resp.GetBody())
		}).
		OnError(func(resp request.IResponse) {
			fmt.Println("error:", resp.GetBody())
		})
}
