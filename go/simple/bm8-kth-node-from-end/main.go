package main

/*
题目：链表中倒数最后 k 个结点
牛客题目标识：BM8
难度：简单
题目链接：https://www.nowcoder.com/practice/886370fe658f41b498d40fb34ae76ff9

输入一个长度为 n 的链表，设链表中的元素值为 ai，返回该链表中倒数第 k 个节点。
如果链表长度小于 k，请返回一个长度为 0 的链表，也就是返回 nil。

输入格式：
输入一个单链表头节点 head 和整数 k。

输出格式：
返回链表中倒数第 k 个节点。判题系统会从该节点开始打印后面的所有节点。

示例 1：
输入：{1,2,3,4,5}, 2
输出：{4,5}
说明：倒数第 2 个节点的值为 4。

示例 2：
输入：{2}, 8
输出：{}
说明：链表长度小于 k，返回空链表。

示例 3：
输入：{1,2,3}, 3
输出：{1,2,3}
说明：倒数第 3 个节点就是头节点。

数据范围：
- 0 <= n <= 10^5
- 0 <= k <= 10^9
- 链表节点值为整数。

要求：
- 时间复杂度 O(n)

进阶：
- 额外空间复杂度 O(1)

函数签名：
func findKthToTail(head *ListNode, k int) *ListNode
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func findKthToTail(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
