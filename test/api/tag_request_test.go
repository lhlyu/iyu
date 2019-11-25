package api

import (
	"fmt"
	"github.com/lhlyu/request"
	"testing"
)

var rq request.IRequest

func init() {
	rq = request.NewRequest()
	rq.SetBaseUrl("http://localhost:8080/api")
}

func TestGetAllTag(t *testing.T) {
	rq.SetUrl("/tag").
		SetMethod(request.GET).
		DoHttp().OnSuccess(func(resp request.IResponse) {
		fmt.Println("success:", resp.GetBody())
	}).OnError(func(resp request.IResponse) {
		fmt.Println("error:", resp.GetBody())
	})
}

func TestInsertTag(t *testing.T) {
	rq.SetUrl("/tag").
		SetData("name:WAZ").
		SetMethod(request.POST).
		DoHttp().OnSuccess(func(resp request.IResponse) {
		fmt.Println("success:", resp.GetBody())
	}).OnError(func(resp request.IResponse) {
		fmt.Println("error:", resp.GetBody())
	})
}

func TestUpdateTag(t *testing.T) {
	rq.SetUrl("/tag").
		SetData(`{"id":15,"name":"8888"}`).
		SetMethod(request.PUT).
		DoHttp().OnSuccess(func(resp request.IResponse) {
		fmt.Println("success:", resp.GetBody())
	}).OnError(func(resp request.IResponse) {
		fmt.Println("error:", resp.GetBody())
	})
}

func TestDeleteTag(t *testing.T) {
	rq.SetUrl("/tag").
		SetData(`{"id":16,"real":1}`).
		SetMethod(request.DELETE).
		DoHttp().OnSuccess(func(resp request.IResponse) {
		fmt.Println("success:", resp.GetBody())
	}).OnError(func(resp request.IResponse) {
		fmt.Println("error:", resp.GetBody())
	})
}
