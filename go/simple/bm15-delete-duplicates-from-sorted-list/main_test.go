package main

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "示例一", input: []int{1, 1, 2}, want: []int{1, 2}},
		{name: "多组重复元素", input: []int{1, 1, 2, 3, 3}, want: []int{1, 2, 3}},
		{name: "空链表", input: nil, want: nil},
		{name: "单个节点", input: []int{5}, want: []int{5}},
		{name: "全部节点重复", input: []int{2, 2, 2, 2}, want: []int{2}},
		{name: "包含负数", input: []int{-3, -3, -1, 0, 0, 2}, want: []int{-3, -1, 0, 2}},
		{name: "没有重复元素", input: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			got := listToSlice(deleteDuplicates(head))
			if !sameInts(got, tt.want) {
				t.Errorf("deleteDuplicates(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func buildList(values []int) *ListNode {
	var dummy ListNode
	current := &dummy
	for _, value := range values {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	values := make([]int, 0)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}

func sameInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
