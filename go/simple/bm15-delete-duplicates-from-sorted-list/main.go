package main

/*
题目：删除有序链表中重复的元素-I
牛客题目标识：BM15
难度：简单
题目链接：https://www.nowcoder.com/practice/c087914fae584da886a0091e877f2c79

删除给出链表中的重复元素（链表中元素从小到大有序），使链表中的所有元素都只出现一次。

例如：
给出的链表为 1 -> 1 -> 2，返回 1 -> 2。
给出的链表为 1 -> 1 -> 2 -> 3 -> 3，返回 1 -> 2 -> 3。

输入格式：
输入一个按从小到大排列的单链表头节点 head。

输出格式：
返回删除重复元素后的单链表头节点。

示例 1：
输入：{1,1,2}
输出：{1,2}

示例 2：
输入：{1,1,2,3,3}
输出：{1,2,3}

示例 3：
输入：{}
输出：{}

数据范围：
- 0 <= n <= 100
- |val| <= 100
- 链表中的元素已经按照从小到大的顺序排列。

进阶：
- 时间复杂度 O(n)
- 额外空间复杂度 O(1)

函数签名：
func deleteDuplicates(head *ListNode) *ListNode
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	res := head
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return res
}
