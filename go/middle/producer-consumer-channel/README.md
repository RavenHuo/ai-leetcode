# 使用 channel 实现生产者-消费者模型

- 题目标识：producer-consumer-channel
- 难度：中等
- 牛客网链接：[Go 高并发编程高频面试题](https://www.nowcoder.com/discuss/759997003932639232)

## 题目描述

请使用 Go 的 goroutine 和 channel 实现一个生产者-消费者模型。生产者按照输入顺序发送数据，消费者接收并返回数据，要求结果顺序与输入一致，并正确处理 channel 的关闭。

## 要求

1. 生产者和消费者运行在独立的 goroutine 中。
2. 生产者发送完成后关闭 channel。
3. 消费者使用 `value, ok := <-ch` 正确判断 channel 是否关闭。
4. 不发生 goroutine 泄漏、死锁或向已关闭 channel 发送数据。

## 函数签名

```go
func producerConsumer(items []int, bufferSize int) []int
```

## 示例

```text
输入：items = [1, 2, 3, 4, 5], bufferSize = 2
输出：[1, 2, 3, 4, 5]

输入：items = [], bufferSize = 1
输出：[]

输入：items = [0, -1, 0, 2], bufferSize = 0
输出：[0, -1, 0, 2]
```

## 解题提示

生产者负责关闭 channel；消费者通过接收表达式返回的 `ok` 判断 channel 是否已关闭。无缓冲 channel 会让发送和接收直接同步，有缓冲 channel 则允许暂存一定数量的数据。
