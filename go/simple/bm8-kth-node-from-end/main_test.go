package main

import "testing"

func TestFindKthToTail(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  []int
	}{
		{name: "示例一倒数第二个", input: []int{1, 2, 3, 4, 5}, k: 2, want: []int{4, 5}},
		{name: "示例二超出链表长度", input: []int{2}, k: 8, want: nil},
		{name: "示例三倒数最后一个", input: []int{1, 2, 3}, k: 3, want: []int{1, 2, 3}},
		{name: "倒数第一个", input: []int{1, 2, 3, 4}, k: 1, want: []int{4}},
		{name: "空链表", input: nil, k: 1, want: nil},
		{name: "无效的零值 k", input: []int{1, 2, 3}, k: 0, want: nil},
		{name: "包含负数", input: []int{-3, -2, -1, 0}, k: 2, want: []int{-1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			got := listToSlice(findKthToTail(head, tt.k))
			if !sameInts(got, tt.want) {
				t.Errorf("findKthToTail(%v, %d) = %v, want %v", tt.input, tt.k, got, tt.want)
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
