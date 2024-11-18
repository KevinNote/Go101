# Basic HTTP Server

[Gin](https://github.com/gin-gonic/gin) 是 Go 语言非常流行的 HTTP 服务器，请使用 Gin 实现简单的 HTTP Server。API 定义如下：

```
[GET] /currentTime
响应描述: 当前机器的时间
响应样例: 2024-11-18 12:00
响应类型: Text

[POST] /add
描述：输入**整数类型**变量 a 和 b，响应一个 sum 请求。
请求样例: { "a": 12, "b": 18 }
请求样例: JSON
响应样例: { "sum": 30 }
响应类型: JSON
提示：可以使用 Gin 自带的 JSON 序列和反序列化机制

[POST] /addString
描述：输入**字符串类型**变量 a 和 b，响应一个 sum 请求。
请求样例: { "a": "12", "b": "18" }
请求样例: JSON
响应样例: { "sum": 30 }
响应类型: JSON
```

你需要让 Gin 监听 `127.0.0.1:11451`。
