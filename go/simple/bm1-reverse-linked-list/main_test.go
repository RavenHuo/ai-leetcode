package main

import "testing"

func TestReverseList(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "示例一", input: []int{1, 2, 3}, want: []int{3, 2, 1}},
		{name: "空链表", input: nil, want: nil},
		{name: "单节点", input: []int{1}, want: []int{1}},
		{name: "两个节点", input: []int{1, 2}, want: []int{2, 1}},
		{name: "包含负数", input: []int{-3, -2, -1, 0}, want: []int{0, -1, -2, -3}},
		{name: "重复值", input: []int{2, 2, 3, 3}, want: []int{3, 3, 2, 2}},
		{name: "较长链表", input: []int{1, 2, 3, 4, 5, 6}, want: []int{6, 5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			got := listToSlice(reverseList(head))
			if !sameInts(got, tt.want) {
				t.Errorf("reverseList(%v) = %v, want %v", tt.input, got, tt.want)
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
