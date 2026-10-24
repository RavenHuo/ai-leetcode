package main

import (
	"sync"
	"time"
)

/*
题目：实现一个线程安全的令牌桶限流器
牛客网题目标识：token-bucket-limiter
难度：中等
牛客网题目链接：https://www.nowcoder.com/discuss/924737812140666880

请使用 Go 实现一个简单的令牌桶限流器。

令牌桶以固定速率生成令牌，桶的容量有限。每次请求到来时需要消耗一个令牌：
- 如果桶中有令牌，则允许请求，并消耗一个令牌；
- 如果桶中没有令牌，则拒绝请求；
- 桶中令牌数量不能超过最大容量。

要求：
1. 支持设置桶容量 capacity 和每秒生成的令牌数 rate。
2. Allow 方法返回当前请求是否被允许。
3. 令牌数量应根据请求到达时间自动补充。
4. 令牌数量不能超过桶容量。
5. 多个 Goroutine 并发调用 Allow 时，不能出现数据竞争。
6. 为了方便测试，Allow 接收请求发生时间 now，而不是在函数内部直接调用 time.Now。

示例 1：
桶容量为 2，初始有 2 个令牌，连续请求 3 次，结果为 true、true、false。

示例 2：
桶容量为 2，速率为每秒 1 个令牌，消耗完令牌后等待 1 秒，再次请求应该被允许。

示例 3：
多个 Goroutine 并发调用 Allow 时，仍然只能按照实际令牌数量允许请求。

提示：
- 使用互斥锁保护令牌数量和上次补充时间。
- 根据 now 与 lastTime 的时间差计算新增令牌数量。
- 可以使用浮点数保存令牌数量，避免丢失小数时间产生的令牌。
- 计算完成后，将 lastTime 更新为 now。
*/

type TokenBucket struct {
	capacity, rate, len float64
	lock                sync.Mutex
	createTime          time.Time
}

func NewTokenBucket(capacity, rate float64, now time.Time) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		rate:       rate,
		len:        capacity,
		createTime: now,
	}
}

func (b *TokenBucket) Allow(now time.Time) bool {
	b.lock.Lock()
	defer b.lock.Unlock()
	s := now.Sub(b.createTime).Seconds()
	create := s * b.rate
	if b.len+create < b.capacity {
		b.len = b.len + create
	} else {
		b.len = b.capacity
	}
	b.createTime = now
	if b.len-1 >= 0 {
		b.len = b.len - 1
		return true
	}
	return false
}
