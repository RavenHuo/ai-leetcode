package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "示例一", input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		{name: "示例二单个元素", input: []int{1}, want: 1},
		{name: "示例三", input: []int{5, 4, -1, 7, 8}, want: 23},
		{name: "全部为负数", input: []int{-3, -1, -2}, want: -1},
		{name: "包含零", input: []int{-2, 0, -1}, want: 0},
		{name: "最大子数组在开头", input: []int{4, -1, 2, -10, 1}, want: 5},
		{name: "最大子数组在末尾", input: []int{-10, 1, 2, 3}, want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.input); got != tt.want {
				t.Errorf("maxSubArray(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
