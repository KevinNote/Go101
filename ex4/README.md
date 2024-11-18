# 锁

Go 语言支持并发，但是并发编程中会遇到一些问题，比如数据竞争。为了解决这个问题，Go 语言提供了一些同步原语，比如 `sync.Mutex`（互斥锁） 和 读写锁 `sync.RWMutex`。

请使用 `sync.Mutex` 实现一个线程安全的计数器 `MutexCounter`，包含 `Add` 和 `Get` 两个方法。

请使用 `sync.RWMutex` 实现一个线程安全的计数器 `RWMutexCounter`，包含 `Add` 和 `Get` 两个方法。

请提交一份不超过150字的报告，说明这两种锁的区别。

Hint: 你可以使用 `defer` 语法糖来释放锁。