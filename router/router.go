package router

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/context"
)

func SetRouter(app *iris.Application){
    app.Get("/", func(ctx context.Context) {
        ctx.Write([]byte("<a href='lhlyu.com'>lhlyu</a>"))
    })
}
