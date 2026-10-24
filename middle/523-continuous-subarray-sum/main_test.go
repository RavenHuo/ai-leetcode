package main

import "testing"

func TestCheckSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "示例一：存在长度为二的子数组",
			nums: []int{23, 2, 4, 6, 7},
			k:    6,
			want: true,
		},
		{
			name: "示例二：整个数组的和是倍数",
			nums: []int{23, 2, 6, 4, 7},
			k:    6,
			want: true,
		},
		{
			name: "示例三：不存在符合条件的子数组",
			nums: []int{23, 2, 6, 4, 7},
			k:    13,
			want: false,
		},
		{
			name: "两个零可以组成和为零的子数组",
			nums: []int{0, 0},
			k:    1,
			want: true,
		},
		{
			name: "连续零和前面的元素",
			nums: []int{1, 0, 0},
			k:    1,
			want: true,
		},
		{
			name: "单个元素不能作为答案",
			nums: []int{5},
			k:    5,
			want: false,
		},
		{
			name: "长度为二但和不是倍数",
			nums: []int{1, 2},
			k:    4,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.nums, tt.k); got != tt.want {
				t.Fatalf("checkSubarraySum(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
