package main

/*
题目：判断链表中是否有环
牛客题目标识：BM6
难度：简单
题目链接：https://www.nowcoder.com/practice/650474f313294468a4ded3ce0f7898b9

判断给定的链表中是否有环。如果有环则返回 true，否则返回 false。

数据范围：
- 0 <= n <= 10000
- |val| <= 100000

要求：
- 时间复杂度 O(n)
- 额外空间复杂度 O(1)

输入说明：
输入分为两部分，第一部分为链表，第二部分代表是否有环，然后将组成的 head 头节点传入函数。
-1 代表无环，其它数字代表有环。这些参数仅用于方便自测，实际编程时传入的是链表头节点。

示例 1：
输入：{3,2,0,-4}, 1
输出：true
说明：尾节点 -4 指向下标为 1 的节点 2，组成环形链表。

示例 2：
输入：{1}, -1
输出：false
说明：链表只有一个节点且没有环。

示例 3：
输入：{-1,-7,7,-4,19,6,-9,-5,-2,-5}, 6
输出：true

函数签名：
func hasCycle(head *ListNode) bool
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
