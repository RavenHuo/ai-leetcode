package main

import "testing"

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		k    int
		want []int
	}{
		{name: "two groups and remainder", head: list(1, 2, 3, 4, 5), k: 2, want: []int{2, 1, 4, 3, 5}},
		{name: "one complete group", head: list(1, 2, 3, 4, 5), k: 3, want: []int{3, 2, 1, 4, 5}},
		{name: "k equals one", head: list(1, 2, 3), k: 1, want: []int{1, 2, 3}},
		{name: "k equals length", head: list(1, 2, 3), k: 3, want: []int{3, 2, 1}},
		{name: "empty list", head: nil, k: 1, want: []int{}},
		{name: "single node", head: list(7), k: 1, want: []int{7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := values(reverseKGroup(tt.head, tt.k))
			if len(got) != len(tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func list(values ...int) *ListNode {
	var head, tail *ListNode
	for _, value := range values {
		node := &ListNode{Val: value}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
	}
	return head
}

func values(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
