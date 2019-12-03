## iyu
iris blog server

### 前端预览(未开发完,非最终版)

[链接]( https://lhlyu.github.io/blog/?v=1.0)

### 项目进度

- 开发分支: develop

1. 第一部分完成(2019-12-03)
2. 第二部分内容

```text
1. 评论、审查
2. 全局配置管理
3. 代码优化
```

### 项目目录结构

```
├─cache        - 缓存
├─common       - 公共 
├─conf         - 配置文件
├─controller   - 控制器
│  └─vo        
├─errcode      - 错误码
├─middleware   - 中间件
├─module       - 模块
├─repository   - 数据访问层
│  ├─po
│  └─test
├─router       - 路由
├─service      - 服务
│  └─bo
├─test         - 测试
│  └─api 
└─util         - 工具
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
