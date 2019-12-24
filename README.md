## iyu
iris blog server

### 前端预览(未开发完,非最终版)

[链接]( https://lhlyu.github.io/blog/?v=1.0)

### 项目进度

- 开发分支: alpha

### 项目目录结构

```
├─cache
├─common
├─conf
├─controller
│  └─dto
├─errcode
├─middleware
├─module
├─repository
│  ├─article_repository
│  ├─category_repository
│  ├─po
│  ├─quanta_repository
│  ├─tag_repository
│  └─user_repository
├─router
├─service
│  ├─article_service
│  ├─category_service
│  ├─quanta_service
│  ├─tag_service
│  ├─user_service
│  └─vo
└─util

```

### 运行

- 下载到本地

> git clone https://github.com/lhlyu/iyu.git

下面的操作都在项目根目录

- 创建数据库

> sql文件在 ./lhlyu_blog.sql
> 库名:lhlyu_blog

- 设置配置(手动更改)

> 配置文件在 ./conf/config.yaml

- 设置代理

go1.11+     
> set GOPROXY=https://goproxy.cn

go1.13+ 
> go env -w GOPROXY=https://goproxy.cn,direct

- 运行
> go run main.go
