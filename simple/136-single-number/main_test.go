package main

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "示例一", input: []int{2, 2, 1}, want: 1},
		{name: "示例二", input: []int{4, 1, 2, 1, 2}, want: 4},
		{name: "单个元素", input: []int{1}, want: 1},
		{name: "唯一元素在中间", input: []int{7, 3, 5, 3, 5}, want: 7},
		{name: "包含负数", input: []int{-1, -2, -1}, want: -2},
		{name: "唯一元素为零", input: []int{8, 8, 0}, want: 0},
		{name: "较大数值", input: []int{30000, -30000, 30000}, want: -30000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.input); got != tt.want {
				t.Errorf("singleNumber(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
