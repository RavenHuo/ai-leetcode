package main

import "testing"

func TestEntryNodeOfLoop(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		entry  int
	}{
		{name: "cycle starts in middle", values: []int{1, 2, 3, 4, 5}, entry: 2},
		{name: "cycle starts at head", values: []int{1, 2, 3}, entry: 0},
		{name: "single node cycle", values: []int{7}, entry: 0},
		{name: "no cycle", values: []int{1, 2, 3}, entry: -1},
		{name: "empty list", values: nil, entry: -1},
		{name: "cycle starts at tail", values: []int{4, 5, 6}, entry: 2},
		{name: "duplicate values", values: []int{1, 2, 1, 3}, entry: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head, want *ListNode
			if tt.entry < 0 {
				head = linearList(tt.values...)
			} else {
				head = cyclicList(tt.values, tt.entry)
				want = head
				for i := 0; i < tt.entry; i++ {
					want = want.Next
				}
			}
			if got := entryNodeOfLoop(head); got != want {
				t.Fatalf("got %v, want %v", nodeValue(got), nodeValue(want))
			}
		})
	}
}

func linearList(values ...int) *ListNode {
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

func cyclicList(values []int, entryIndex int) *ListNode {
	head := linearList(values...)
	if head == nil {
		return nil
	}

	var entry, tail *ListNode
	for i, node := 0, head; node != nil; i, node = i+1, node.Next {
		if i == entryIndex {
			entry = node
		}
		tail = node
	}
	tail.Next = entry
	return head
}

func nodeValue(node *ListNode) any {
	if node == nil {
		return nil
	}
	return node.Val
}
