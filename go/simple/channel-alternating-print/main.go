package main

import (
	"strconv"
	"sync"
)

/*
题目：使用两个 goroutine 和 channel 交替打印字母与数字
题目标识：Go-Channel-Alternating-Print
难度：面试题
参考来源：https://www.nowcoder.com/discuss/712384113688199168

请使用两个 goroutine，借助 channel 实现字母和数字的交替输出。
第一个 goroutine 负责输出大写字母，第二个 goroutine 负责输出数字。

当 n = 3 时，期望输出：A1B2C3
当 n = 26 时，期望输出：A1B2C3...Z26

要求：
1. 必须使用两个 goroutine。
2. 必须使用 channel 控制两个 goroutine 的执行顺序。
3. 字母和数字必须严格交替输出，不能出现 A2B1 等顺序错误。
4. 函数必须等待两个 goroutine 都执行结束后再返回。
5. 不允许依赖 sleep 控制执行顺序，也不能通过固定的调度顺序实现。

输入格式：
输入一个整数 n，表示需要输出的字母和数字对数。

输出格式：
返回交替拼接后的字符串。每一对由一个大写字母和一个数字组成。

示例 1：
输入：n = 3
输出："A1B2C3"

示例 2：
输入：n = 1
输出："A1"

示例 3：
输入：n = 0
输出：""

提示：
- n 小于等于 26。
- n <= 0 时返回空字符串。
- 注意处理 channel 的发送、接收和关闭，避免程序死锁。

函数签名：
func alternatePrint(n int) string
*/

func alternatePrint(n int) string {
	if n <= 0 {
		return ""
	}
	res := ""
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			select {
			case a := <-ch1:
				if a > n {
					close(ch2)
					return
				}
				res += string(rune('A' + (a - 1)))
				ch2 <- a
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case a, ok := <-ch2:
				if !ok {
					return
				}
				res += strconv.Itoa(a)
				ch1 <- a + 1
			}
		}
	}()
	ch1 <- 1
	wg.Wait()
	return res
}
