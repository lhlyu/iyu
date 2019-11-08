## api 

1. [首页文章列表](#1-首页文章列表)
2. [单独文章](#2-单独文章)
3. [作者信息](#3-作者信息)
4. [网站信息](#4-网站信息)
5. [分类](#5-分类)
6. [标签](#6-标签)

#### 1. 首页文章列表

- 请求

```text
GET /api/articles?pageNum=1&pageSize=10&category=0&tag=0
```

- 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "page": {
      "pageNum": 1,
      "pageSize": 10,
      "total": 11
    },
    "list": [
      {
        "id": 1,
        "title": "",
        "description": "",
        "date": "",
        "fire": 11,
        "commentNum": 1,
        "wraper": "",
        "nail": {
          "id": 1,
          "name": "",
          "color": ""
        },
        "category": {
          "id": 1,
          "name": ""
        },
        "tags": [
          {
            "id": 1,
            "name": ""
          }
        ]
      }
    ]
  }
}
```

#### 2. 单独文章

- 请求

```text
GET /api/articles/:id
```

- 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "id": 1,
    "title": "",
    "content": "",
    "date": "",
    "fire": 11,
    "commentNum": 1,
    "wraper": "",
    "nail": {
      "id": 1,
      "name": "",
      "color": ""
    },
    "category": {
      "id": 1,
      "name": ""
    },
    "tags": [
      {
        "id": 1,
        "name": ""
      }
    ]
  }
}
```


#### 3. 作者信息

- 请求

```text
GET /api/author/info
```

- 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "id": 1,
    "name": "",
    "bio": "",
    "createdAt": 11111,
    "updatedAt": 122,
    "contact": {
      "github": "",
      "gitee": "",
      "qq": "",
      "email": "",
      "phone": ""
    }
  }
}
```


#### 4. 网站信息

- 请求

```text
GET /api/website/info
```

- 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "bg":""
  }
}
```

#### 5. 分类

- 请求

```text
GET /api/categorys
```

- 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "name": "",
        "number": "+11k"
      }
    ]
  }
}
```

#### 6. 标签

- 请求

```text
GET /api/tags
```

- 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "name": ""
      }
    ]
  }
}
```
