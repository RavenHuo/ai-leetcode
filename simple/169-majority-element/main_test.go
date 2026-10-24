package main

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "示例一", input: []int{3, 2, 3}, want: 3},
		{name: "示例二", input: []int{2, 2, 1, 1, 1, 2, 2}, want: 2},
		{name: "只有一个元素", input: []int{1}, want: 1},
		{name: "负数多数元素", input: []int{-1, -1, 2}, want: -1},
		{name: "多数元素在末尾", input: []int{1, 2, 2, 2}, want: 2},
		{name: "多数元素在开头", input: []int{4, 4, 4, 1, 2}, want: 4},
		{name: "较大数值", input: []int{1000000000, -1, 1000000000}, want: 1000000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.input); got != tt.want {
				t.Errorf("majorityElement(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
