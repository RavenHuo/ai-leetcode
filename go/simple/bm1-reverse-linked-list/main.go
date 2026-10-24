package main

/*
题目：反转链表
牛客题目标识：BM1
难度：简单
题目链接：https://www.nowcoder.com/practice/75e878df47f24fdc9dc3e400ec6058ca

给定一个单链表的头节点 pHead，反转该链表后，返回新链表的表头。

输入格式：
输入一个单链表的头节点 pHead。

输出格式：
返回反转后的单链表头节点。

示例 1：
输入：{1,2,3}
输出：{3,2,1}

示例 2：
输入：{}
输出：{}

示例 3：
输入：{1}
输出：{1}

数据范围：
- 0 <= n <= 1000
- 链表节点值为整数。

要求：
- 时间复杂度 O(n)
- 额外空间复杂度 O(1)

函数签名：
func reverseList(head *ListNode) *ListNode
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
