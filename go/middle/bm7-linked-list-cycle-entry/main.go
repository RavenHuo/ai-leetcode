package main

/*
题目：链表中环的入口结点
牛客网题目标识：BM7
难度：中等
题目链接：https://www.nowcoder.com/practice/253d2c59ec3e4bc68da16833f79a38e4

给一个长度为 n 的链表，若其中包含环，请找出链表环的入口结点；否则返回 nil。

输入描述：输入分为入环前的链表部分和环的部分，后台会根据环部分是否为空组装链表。
输出描述：返回环的入口结点；如果没有环，返回 nil。

数据范围：0 <= n <= 100000，要求空间复杂度 O(1)，时间复杂度 O(n)。

示例 1：入环前 {1,2}，环部分 {3,4,5}，返回值为 3。
示例 2：入环前 {1}，环部分 {}，返回 nil。
示例 3：入环前 {}，环部分 {2}，返回值为 2。

解题提示：先使用快慢指针判断是否相遇；相遇后，将一个指针移回头节点，两个指针同步移动，它们再次相遇的位置就是环入口。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func entryNodeOfLoop(head *ListNode) *ListNode {
	var fast, slow *ListNode
	if head == nil || head.Next == nil {
		return nil
	}

	fast = head
	slow = head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}

	}

	return nil
}
