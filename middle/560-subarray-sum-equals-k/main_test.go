package main

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{name: "重复元素", nums: []int{1, 1, 1}, k: 2, want: 2},
		{name: "多个连续子数组", nums: []int{1, 2, 3}, k: 3, want: 2},
		{name: "包含负数", nums: []int{1, -1, 0}, k: 0, want: 3},
		{name: "全为零", nums: []int{0, 0, 0, 0}, k: 0, want: 10},
		{name: "不存在目标和", nums: []int{1, 2, 3}, k: 10, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.nums, tt.k); got != tt.want {
				t.Fatalf("subarraySum(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
