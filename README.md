1.项目的运行环境：
golang 语言支持的运行库 
2.启动方式：
通过 main.go 启动

3.postman 接口进行测试数据
localhost:8080/register 用户注册
request:
{
"name": "jxd",
"username": "jxd111111",
"password": "123456"
}
response:
{
"message": "User registered successfully"
}

localhost:8080/login 登陆
request:
{
"username": "jxd111111",
"password": "123456"
}
response:
{
"message": "User login successfully"
}

localhost:8080/user/createPost
创建文章： 未登录状态
request:
{
"title": "title-测试-新增5",
"data": "asdasdasdsfsdfdsfsdf"
}
response:
{
"error": "未登录"
}

创建文章： 校验
request:
{
"data": "asdasdasdsfsdfdsfsdf"
}
response:
{
"error": "文章内容或者 标题不能为空"
}
创建文章：成功
request:
{
"title": "title-测试-新增5",
"data": "asdasdasdsfsdfdsfsdf"
}
response:
{
"code": "sucess",
"message": "创建文章成功"
}
localhost:8080/allPost 获取所有文章信息
response:
[
{
"Id": 1,
"AuthorId": 1,
"Title": "信审",
"Data": "信审数据12312312",
"CommentState": "",
"CreateTime": "2025-05-30T15:39:25+08:00",
"UpdateTime": "2025-05-30T15:39:27+08:00",
"Comments": null
},
{
"Id": 2,
"AuthorId": 2,
"Title": "请款",
"Data": "请款数据12312312",
"CommentState": "无评论",
"CreateTime": "2025-05-30T15:39:25+08:00",
"UpdateTime": "2025-05-30T15:39:27+08:00",
"Comments": null
},
{
"Id": 9,
"AuthorId": 1,
"Title": "新增内容标题",
"Data": "新增内容1111",
"CommentState": "",
"CreateTime": "2025-05-31T15:03:12+08:00",
"UpdateTime": "2025-05-31T15:03:12+08:00",
"Comments": null
},
{
"Id": 21,
"AuthorId": 2,
"Title": "title-测试-更新4",
"Data": "data-内容sasdsdfsdfsdsdfdsf 4444",
"CommentState": "",
"CreateTime": "2025-06-02T11:12:34+08:00",
"UpdateTime": "2025-06-02T15:42:23+08:00",
"Comments": null
},
{
"Id": 22,
"AuthorId": 0,
"Title": "title-测试-新增5",
"Data": "asdasdasdsfsdfdsfsdf",
"CommentState": "",
"CreateTime": "2025-06-03T14:37:53+08:00",
"UpdateTime": "2025-06-03T14:37:53+08:00",
"Comments": null
}
]


localhost:8080/user/updatePost 更新文章
失败
request:
{
"id": 21,
"title": "title-测试-新增5",
"data": "asdasdasdsfsdfdsfsdf"
}
response:
{
"error": "userID not Post in id 该文章不是作者的"
}
成功
request:
{
"id": 25,
"title": "title-测试-新增5",
"data": "asdasdasdsfsdfdsfsdf"
}
response:
{
"code": "sucess",
"message": "更新文章成功"
}
localhost:8080/user/deletePost 删除文章
失败
request:
{
"id": 7
}
response:
{
"error": "userID not Post in id 该文章不是作者的"
}
成功

request:
{
"id": 25
}
response:
{
"code": "sucess",
"message": "删除文章成功"
}

localhost:8080/user/createCommon 发表评论
失败
request:
{
"postId": 25,
"remark": "评论测试asdasdsfsdfdsfsdf"
}
response:
{
"error": "Invalid token"
}
成功
request:
{
"postId": 25,
"remake": "评论测试asdasdsfsdfdsfsdf"
}
response:
{
"code": "sucess",
"message": "评论成功"
}

GET : localhost:8080/allComment 获取某篇文章评论
request:
{
"postId": 25
}
response:
[
{
"Id": 7,
"AuthorId": 7,
"PostId": 25,
"Remake": "评论测试asdasdsfsdfdsfsdf",
"CreateTime": "2025-06-03T14:57:47+08:00",
"UpdateTime": "2025-06-03T14:57:47+08:00"
},
{
"Id": 8,
"AuthorId": 7,
"PostId": 25,
"Remake": "评论测试2222asdasdsfsdfdsfsdf",
"CreateTime": "2025-06-03T14:57:47+08:00",
"UpdateTime": "2025-06-03T14:57:47+08:00"
}
]
