package main

import "testing"

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "中心在中间", nums: []int{1, 7, 3, 6, 5, 6}, want: 3},
		{name: "不存在中心下标", nums: []int{1, 2, 3}, want: -1},
		{name: "中心在最左边", nums: []int{2, 1, -1}, want: 0},
		{name: "中心在最右边", nums: []int{-1, 1, 2}, want: 2},
		{name: "多个中心返回最左边", nums: []int{0, 0, 0}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.nums); got != tt.want {
				t.Fatalf("pivotIndex(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
