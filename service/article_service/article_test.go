package article_service

import (
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/yutil"
	"testing"
)

func init() {
	module.Register(module.CfgModule, // 读取配置 <必须>
		module.LgModule, // 日志
		module.DbModule, // 连接数据库
		module.InitiateModule,
		module.RedisModule) // redis
	module.Init()
}

func TestService_GetArticleById(t *testing.T) {
	svc := NewService("test")
	result := svc.GetArticleById(2)
	t.Log(yutil.JsonObjToStr(result))
}

func TestService_GetTimeline(t *testing.T) {
	svc := NewService("test")
	result := svc.GetTimeline()
	t.Log(yutil.JsonObjToStr(result))
}

func TestService_AddArticle(t *testing.T) {

	content := `
### docker建立桥接

- 建立一个叫"my-bridge"的网络桥接
> docker network create -d bridge my-bridge

- 查看所有的网桥
> docker network ls

- 查看网桥内所有关联的容器
> docker inspect my-bridge

- 网桥添加一个容器
> docker network connect my-bridage xx-container

- 启动时添加到网桥
> docker run --network my-bridge -p 8080:80 -d nginx:latest

`

	param := &dto.ArticleEditDto{
		UserId:     1,
		Wrapper:    "",
		Title:      "docker建立桥接",
		Summary:    "",
		Content:    content,
		IsTop:      1,
		CategoryId: 1,
		Kind:       1,
		SortNum:    0,
		CmntStatus: 1,
		IsDelete:   1,
		TagIds:     []int{1, 2, 3},
	}
	svc := NewService("test")
	result := svc.AddArticle(param)
	t.Log(yutil.JsonObjToStr(result))
}
