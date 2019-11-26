package api

import (
	"fmt"
	"github.com/lhlyu/request"
	"testing"
)

func TestGetAllQuanta(t *testing.T) {
	rq.SetUrl("/quanta").
		SetMethod(request.GET).
		SetParam("pageNum=1&pageSize=10").
		DoHttp().
		OnSuccess(func(resp request.IResponse) {
			fmt.Println("success:", resp.GetBody())
		}).
		OnError(func(resp request.IResponse) {
			fmt.Println("error:", resp.GetBody())
		})
}

func TestInsertQuanta(t *testing.T) {
	rq.SetUrl("/quanta").
		SetMethod(request.POST).
		SetData(`{"key":"admin.test","value":"1"}`).
		DoHttp().
		OnSuccess(func(resp request.IResponse) {
			fmt.Println("success:", resp.GetBody())
		}).
		OnError(func(resp request.IResponse) {
			fmt.Println("error:", resp.GetBody())
		})
}

func TestUpdateQuanta(t *testing.T) {
	rq.SetUrl("/quanta").
		SetMethod(request.PUT).
		SetData(`{"id":3,"key":"admin.pass","value":"3"}`).
		DoHttp().
		OnSuccess(func(resp request.IResponse) {
			fmt.Println("success:", resp.GetBody())
		}).
		OnError(func(resp request.IResponse) {
			fmt.Println("error:", resp.GetBody())
		})
}

func TestDeleteQuanta(t *testing.T) {
	rq.SetUrl("/quanta").
		SetMethod(request.DELETE).
		SetData(`{"id":3,"real":1}`).
		DoHttp().
		OnSuccess(func(resp request.IResponse) {
			fmt.Println("success:", resp.GetBody())
		}).
		OnError(func(resp request.IResponse) {
			fmt.Println("error:", resp.GetBody())
		})
}
