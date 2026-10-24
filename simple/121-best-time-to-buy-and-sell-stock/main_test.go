package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "典型情况",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "价格持续下降",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "先涨后跌",
			prices: []int{2, 4, 1},
			want:   2,
		},
		{
			name:   "只有一天",
			prices: []int{5},
			want:   0,
		},
		{
			name:   "最低点出现在中间",
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			want:   4,
		},
		{
			name:   "两天上涨",
			prices: []int{1, 2},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
