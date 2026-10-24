# 实现一个简单的 Go Worker Pool

- 牛客网题目标识：`worker-pool`
- 难度：中等
- 牛客网题目链接：https://www.nowcoder.com/discuss/842908696819593216

请使用 Goroutine 和 Channel 实现一个简单的 Worker Pool。

给定一组整数任务 `tasks` 和 Worker 数量 `workerCount`，每个 Worker 负责从任务队列中取出一个整数，将它转换为该整数的平方后返回。函数需要等待所有任务处理完成，并按照任务在 `tasks` 中的原始顺序返回结果。

## 要求

1. 创建 `workerCount` 个 Worker Goroutine。
2. 多个 Worker 通过同一个任务 Channel 获取任务。
3. 所有任务都处理完成后，函数才能返回。
4. 返回结果的顺序必须与 `tasks` 的输入顺序一致。
5. `workerCount <= 0` 时应该安全处理，不能导致程序死锁。

## 示例

### 示例 1

输入：`tasks = [1, 2, 3, 4]`，`workerCount = 2`

输出：`[1, 4, 9, 16]`

### 示例 2

输入：`tasks = [-2, 0, 5]`，`workerCount = 3`

输出：`[4, 0, 25]`

### 示例 3

输入：`tasks = []`，`workerCount = 2`

输出：`[]`

## 解题提示

可以使用任务 Channel 分发任务，使用结果 Channel 收集结果。每个任务需要携带原始下标，避免并发追加结果造成数据竞争，并保证返回顺序。使用 `WaitGroup` 等待所有 Worker 退出，并由合适的一方关闭结果 Channel。
