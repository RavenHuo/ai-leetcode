package main

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			name:         "示例一",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "示例二：连续升温",
			temperatures: []int{30, 40, 50, 60},
			want:         []int{1, 1, 1, 0},
		},
		{
			name:         "示例三",
			temperatures: []int{30, 60, 90},
			want:         []int{1, 1, 0},
		},
		{
			name:         "连续降温",
			temperatures: []int{30, 29, 28, 27},
			want:         []int{0, 0, 0, 0},
		},
		{
			name:         "温度相同不算升高",
			temperatures: []int{70, 70, 70},
			want:         []int{0, 0, 0},
		},
		{
			name:         "单个温度",
			temperatures: []int{100},
			want:         []int{0},
		},
		{
			name:         "中间温度先遇到更高温度",
			temperatures: []int{73, 72, 75},
			want:         []int{2, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dailyTemperatures(tt.temperatures)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("dailyTemperatures(%v) = %v, want %v", tt.temperatures, got, tt.want)
			}
		})
	}
}
