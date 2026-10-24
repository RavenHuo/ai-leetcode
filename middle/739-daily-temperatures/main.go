package main

/*
题目：每日温度
LeetCode 题号：739
难度：中等
题目链接：https://leetcode.cn/problems/daily-temperatures/description/

给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，
下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

示例 1：

输入：temperatures = [73,74,75,71,69,72,76,73]
输出：[1,1,4,2,1,1,0,0]

示例 2：

输入：temperatures = [30,40,50,60]
输出：[1,1,1,0]

示例 3：

输入：temperatures = [30,60,90]
输出：[1,1,0]

提示：

    1 <= temperatures.length <= 10^5
    30 <= temperatures[i] <= 100

解题提示：
从左到右维护一个存放下标的单调递减栈。当当前温度高于栈顶下标对应的温度时，
当前日期就是栈顶日期等待的下一个更高温度，弹出并计算下标差；处理完后再将当前下标入栈。

函数签名：
func dailyTemperatures(temperatures []int) []int
*/

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0) // 存下标
	res := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		// 判断如果栈顶是小于当前值时，出栈
		for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// 出栈
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[pre] = i - pre
		}
		stack = append(stack, i)
	}
	return res
}
