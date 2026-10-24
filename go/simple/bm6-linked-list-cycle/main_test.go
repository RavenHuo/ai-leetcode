package main

import "testing"

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{name: "空链表", head: nil, want: false},
		{name: "单节点无环", head: buildCycleList([]int{1}, -1), want: false},
		{name: "单节点自环", head: buildCycleList([]int{1}, 0), want: true},
		{name: "尾节点连接中间节点", head: buildCycleList([]int{3, 2, 0, -4}, 1), want: true},
		{name: "普通无环链表", head: buildCycleList([]int{1, 2, 3, 4}, -1), want: false},
		{name: "尾节点连接头节点", head: buildCycleList([]int{1, 2, 3}, 0), want: true},
		{name: "较长无环链表", head: buildCycleList([]int{-1, -7, 7, -4, 19, 6}, -1), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildCycleList(values []int, entry int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(values))
	for i, value := range values {
		nodes[i] = &ListNode{Val: value}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	if entry >= 0 && entry < len(nodes) {
		nodes[len(nodes)-1].Next = nodes[entry]
	}
	return nodes[0]
}
