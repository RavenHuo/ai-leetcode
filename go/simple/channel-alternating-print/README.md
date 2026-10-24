# 使用两个 goroutine 和 channel 交替打印字母与数字

- 题目标识：`Go-Channel-Alternating-Print`
- 难度：面试题
- 参考来源：[Go语言高频面试题—两个协程交替打印 - 牛客网](https://www.nowcoder.com/discuss/712384113688199168)

## 题目描述

请使用两个 goroutine，借助 channel 实现字母和数字的交替输出。

第一个 goroutine 负责输出大写字母，第二个 goroutine 负责输出数字。当 `n = 3` 时，期望输出 `A1B2C3`；当 `n = 26` 时，期望输出 `A1B2C3...Z26`。

## 要求

1. 必须使用两个 goroutine。
2. 必须使用 channel 控制两个 goroutine 的执行顺序。
3. 字母和数字必须严格交替输出。
4. 函数必须等待两个 goroutine 都执行结束后再返回。
5. 不允许依赖 `sleep` 控制执行顺序。

## 输入格式

输入一个整数 `n`，表示需要输出的字母和数字对数。

## 输出格式

返回交替拼接后的字符串，每一对由一个大写字母和一个数字组成。

## 数据范围

- `n <= 26`
- `n <= 0` 时返回空字符串。

## 示例

### 示例 1

输入：`n = 3`

输出：`"A1B2C3"`

### 示例 2

输入：`n = 1`

输出：`"A1"`

### 示例 3

输入：`n = 0`

输出：`""`

## 解题提示

可以创建两个无缓冲 channel：一个通知字母 goroutine，一个通知数字 goroutine。字母 goroutine 输出一个字母后通知数字 goroutine，数字 goroutine 输出一个数字后再通知字母 goroutine。最后使用 `sync.WaitGroup` 等待两个 goroutine 完成。
