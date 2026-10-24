package main

/*
题目：最小栈
LeetCode 题号：155
难度：中等
题目链接：https://leetcode.cn/problems/min-stack/

设计一个支持 push、pop、top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类：
- MinStack() 初始化堆栈对象。
- void push(int val) 将元素 val 推入堆栈。
- void pop() 删除堆栈顶部的元素。
- int top() 获取堆栈顶部元素。
- int getMin() 获取堆栈中的最小元素。

要求：每个函数的时间复杂度必须是 O(1)。

示例 1：
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]
输出：
[null,null,null,null,-3,null,0,-2]

示例 2：
输入：
["MinStack","push","getMin","top","pop","getMin"]
[[],[5],[],[],[],[]]
输出：
[null,null,5,5,null,0]

提示：
- -2^31 <= val <= 2^31 - 1
- pop、top 和 getMin 操作调用总次数不超过 3 * 10^4 次。
- top、pop 和 getMin 操作总是在非空栈上调用。

解题提示：
除了保存普通元素外，再维护一个辅助栈记录每个时刻的最小值。
每次 push 时同步更新最小值，pop 时两个栈一起弹出，即可保证 getMin 为 O(1)。
*/

type MinStack struct {
	// TODO: 请定义保存元素和最小值所需的数据结构
}

func Constructor() MinStack {
	// TODO: 请完成初始化
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	// TODO: 请完成入栈
}

func (s *MinStack) Pop() {
	// TODO: 请完成出栈
}

func (s *MinStack) Top() int {
	// TODO: 请返回栈顶元素
	return 0
}

func (s *MinStack) GetMin() int {
	// TODO: 请返回当前最小值
	return 0
}
