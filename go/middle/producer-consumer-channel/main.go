package main

import "sync"

/*
题目：使用 channel 实现生产者-消费者模型
难度：中等
题目标识：producer-consumer-channel
题目链接：https://www.nowcoder.com/discuss/759997003932639232

题目描述：
请使用 Go 的 goroutine 和 channel 实现一个生产者-消费者模型。
生产者负责按照输入顺序将 items 中的元素发送到 channel，消费者负责从
channel 中接收元素并返回。channel 的缓冲区大小由 bufferSize 指定。

要求：
1. 生产者和消费者必须运行在独立的 goroutine 中。
2. 生产者发送完所有元素后，必须正确关闭 channel。
3. 消费者必须正确判断 channel 是否已经关闭，不能因为关闭后的零值而多消费数据。
4. 返回结果中的元素顺序必须与 items 一致。
5. 不允许发生 goroutine 泄漏、死锁或向已关闭的 channel 发送数据。

输入格式：
items 是待生产的数据；bufferSize 是 channel 的缓冲区大小。

输出格式：
返回消费者按照接收顺序收集到的全部数据。

示例 1：
输入：items = [1, 2, 3, 4, 5], bufferSize = 2
输出：[1, 2, 3, 4, 5]

示例 2：
输入：items = [], bufferSize = 1
输出：[]

示例 3：
输入：items = [0, -1, 0, 2], bufferSize = 0
输出：[0, -1, 0, 2]

提示：
- 生产者通常应该负责关闭 channel，因为生产者知道什么时候不会再发送数据。
- 接收 channel 时可以使用 value, ok := <-ch 判断 channel 是否关闭。
- 注意无缓冲 channel 和有缓冲 channel 的行为差异。
*/

// 解题提示：先创建指定容量的 channel，再分别启动生产者和消费者 goroutine；
// 生产者发送完成后关闭 channel，消费者通过 ok 判断接收是否结束。
func producerConsumer(items []int, bufferSize int) []int {
	res := make([]int, 0)
	ch1 := make(chan int, bufferSize)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < len(items); i++ {
			ch1 <- items[i]
		}
		close(ch1)

	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case a, ok := <-ch1:
				if !ok {
					return
				}
				res = append(res, a)
			}
		}
	}()
	wg.Wait()
	return res
}
