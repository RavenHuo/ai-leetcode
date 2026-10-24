package main

import (
	"sync"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	base := time.Unix(0, 0)

	t.Run("桶容量限制", func(t *testing.T) {
		bucket := NewTokenBucket(2, 1, base)
		if !bucket.Allow(base) || !bucket.Allow(base) || bucket.Allow(base) {
			t.Fatal("令牌数量不符合桶容量限制")
		}
	})

	t.Run("按时间补充令牌", func(t *testing.T) {
		bucket := NewTokenBucket(2, 1, base)
		if !bucket.Allow(base) || !bucket.Allow(base) {
			t.Fatal("初始令牌应该允许前两次请求")
		}
		if bucket.Allow(base.Add(500 * time.Millisecond)) {
			t.Fatal("半秒只能生成半个令牌，不应允许请求")
		}
		if !bucket.Allow(base.Add(time.Second)) {
			t.Fatal("一秒后应该补充一个令牌")
		}
	})

	t.Run("不超过最大容量", func(t *testing.T) {
		bucket := NewTokenBucket(2, 10, base)
		if !bucket.Allow(base) || !bucket.Allow(base) || bucket.Allow(base) {
			t.Fatal("初始令牌数量错误")
		}
		if !bucket.Allow(base.Add(time.Second)) || !bucket.Allow(base.Add(time.Second)) {
			t.Fatal("补充后的令牌数量不正确")
		}
		if bucket.Allow(base.Add(time.Second)) {
			t.Fatal("令牌数量不应超过容量")
		}
	})

	t.Run("时间不倒退", func(t *testing.T) {
		bucket := NewTokenBucket(1, 1, base)
		if !bucket.Allow(base) {
			t.Fatal("第一次请求应该成功")
		}
		if bucket.Allow(base.Add(-time.Second)) {
			t.Fatal("时间倒退时不应凭空产生令牌")
		}
	})

	t.Run("并发访问", func(t *testing.T) {
		bucket := NewTokenBucket(100, 0, base)
		var wg sync.WaitGroup
		var mu sync.Mutex
		allowed := 0
		for i := 0; i < 200; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if bucket.Allow(base) {
					mu.Lock()
					allowed++
					mu.Unlock()
				}
			}()
		}
		wg.Wait()
		if allowed != 100 {
			t.Fatalf("允许请求数 = %d, want 100", allowed)
		}
	})
}
