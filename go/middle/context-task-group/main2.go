package main

import (
	"context"
	"sync"
)

// runTaskGroup2 是 runTaskGroup 的完整实现版本。
// 任务函数需要主动监听 ctx.Done()，才能在取消后及时退出。
func runTaskGroup2(parent context.Context, tasks []func(context.Context) error) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	var wg sync.WaitGroup
	var once sync.Once
	var firstErr error

	// 容量为 1，只保存第一个错误，不让错误上报阻塞任务退出。
	errCh := make(chan error, 1)

	for _, task := range tasks {
		task := task
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := task(ctx); err != nil {
				once.Do(func() {
					firstErr = err
					cancel()
					errCh <- err
				})
			}
		}()
	}

	// 必须等待所有任务退出，避免函数返回后遗留 Goroutine。
	wg.Wait()

	// 读取错误 Channel，保留错误 Channel 的使用语义。
	select {
	case err := <-errCh:
		return err
	default:
	}

	// 如果父 context 被取消，但任务本身没有返回错误，返回父 context 的错误。
	if parent.Err() != nil {
		return parent.Err()
	}

	return firstErr
}
