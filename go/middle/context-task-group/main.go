package main

import (
	"context"
	"errors"
	"sync"
)

/*
题目：实现一个可取消的并发任务组
牛客网题目标识：context-task-group
难度：中等
牛客网题目链接：https://www.nowcoder.com/discuss/927512391112323072

给定一个父 context 和一组任务函数，请并发执行这些任务，并实现以下要求：

1. 所有任务都在独立的 Goroutine 中执行。
2. 任意一个任务返回错误时，立即取消其他仍在执行的任务。
3. 返回第一个发生的错误；如果所有任务成功，返回 nil。
4. 父 context 被取消或超时时，所有任务都应该能够收到取消通知并退出。
5. 不能因为某个任务阻塞而导致任务组永久等待。

任务函数通过 context.Context 感知取消信号。任务内部应该在适当的位置检查
ctx.Done()，并在收到取消通知后及时返回。

示例 1：
输入：3 个都返回 nil 的任务
输出：nil

示例 2：
输入：其中一个任务返回 errors.New("failed")
输出：返回该错误

示例 3：
输入：父 context 超时
输出：返回 context.DeadlineExceeded 或任务返回的相关错误

提示：
- 使用 context.WithCancel 建立子 context。
- 使用 sync.WaitGroup 等待任务结束。
- 使用带缓冲的错误 Channel 保存第一个错误，避免错误发送阻塞。
- 任务完成后必须调用 WaitGroup.Done()，推荐使用 defer。
- 取消只是一种通知，任务函数必须主动监听 ctx.Done() 才能及时退出。
*/

var errTaskFailed = errors.New("task failed")

func runTaskGroup(parent context.Context, tasks []func(context.Context) error) error {
	errChannel := make(chan error, len(tasks))
	c := make(chan struct{})
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(parent)
	for i := 0; i < len(tasks); i++ {
		task := tasks[i]
		wg.Add(1)

		go func() {
			defer wg.Done()
			err := task(ctx)
			if err != nil {
				cancel()
				errChannel <- err
			}
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	for {
		select {
		case e := <-errChannel:
			if e != nil {
				return errTaskFailed
			}
		case <-parent.Done():
			cancel()
			if parent.Err() != nil {
				return parent.Err()
			}
		case <-c:
			return nil
		}
	}

}
