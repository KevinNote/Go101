# 序列化和反序列化

JSON 是常见的数据交换格式，Go 语言内置对 JSON 的支持。JSON 包含了两个函数：`Marshal` 和 `Unmarshal`，用于序列化和反序列化。

请使用 `struct` 定义一个 `Person` 结构体，包含 `name` 和 `age` 两个字段。然后，将 `Person` 结构体实例序列化为 JSON 字符串，并打印出来。

反序列化时，将 JSON 字符串反序列化为 `Person` 结构体实例，并打印出来。

这里你不应该使用 `map` 类型，而是使用 `struct` 类型。

Hint: 在 Go 中，string 可以和 []byte 进行相互转换：`string([]byte)` 和 `[]byte(string)`。