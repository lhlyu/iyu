# iyu
iris blog server

## 项目进度

- 正在开发...
- 开发分支: develop

## 项目目录结构

```
├─cache           - 缓存
├─common          - 公共
├─conf            - 配置文件
├─controller      - 控制器
|  └─vo              
├─errcode         - 错误码
├─middleware      - 中间件
├─module          - 模块
├─repository      - 数据访问
│  └─po           - 数据库持久对象（这个用工具生成，工具地址：github.com/lhlyu/got）
├─router          - 路由
├─test            - 测试
├─util            - 工具
  

```

## 运行

- 下载到本地

> git clone https://github.com/lhlyu/iyu.git

下面的操作都在项目根目录

- 创建数据库

> sql文件在 ./repository/po/lhlyu_blog.sql
> 库名:lhlyu_blog

- 设置配置

> 配置文件在 ./conf/config.yaml

- 设置代理

go1.11+     
> set GOPROXY=https://goproxy.cn

go1.13+ 
> go env -w GOPROXY=https://goproxy.cn,direct

- 运行
> go run main.go
