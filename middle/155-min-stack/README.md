# 155. 最小栈

- 难度：中等
- LeetCode 链接：https://leetcode.cn/problems/min-stack/

设计一个支持 `push`、`pop`、`top` 操作，并能在常数时间内检索到最小元素的栈。

## 操作

- `MinStack()` 初始化堆栈对象。
- `Push(val)` 将元素 `val` 推入堆栈。
- `Pop()` 删除堆栈顶部的元素。
- `Top()` 获取堆栈顶部元素。
- `GetMin()` 获取堆栈中的最小元素。

要求每个函数的时间复杂度必须是 `O(1)`。

## 示例

### 示例 1

操作：`Push(-2), Push(0), Push(-3), GetMin(), Pop(), Top(), GetMin()`

输出：`-3, 0, -2`

### 示例 2

操作：`Push(5), GetMin(), Top()`

输出：`5, 5`

### 示例 3

操作：`Push(2), Push(1), Push(1), Pop(), GetMin()`

输出：`1`

## 数据范围

- `-2^31 <= val <= 2^31 - 1`
- `pop`、`top` 和 `getMin` 操作总次数不超过 `3 * 10^4` 次。
- `top`、`pop` 和 `getMin` 操作总是在非空栈上调用。

## 解题提示

除了保存普通元素外，再维护一个辅助栈记录每个时刻的最小值。每次 `Push` 时同步更新最小值，`Pop` 时两个栈一起弹出，即可保证 `GetMin` 为 `O(1)`。
