package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{name: "多个区间合并", intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{name: "边界相连", intervals: [][]int{{1, 4}, {4, 5}}, want: [][]int{{1, 5}}},
		{name: "完全包含", intervals: [][]int{{1, 10}, {2, 3}}, want: [][]int{{1, 10}}},
		{name: "单个区间", intervals: [][]int{{5, 7}}, want: [][]int{{5, 7}}},
		{name: "无重叠区间乱序输入", intervals: [][]int{{8, 9}, {1, 2}, {4, 5}}, want: [][]int{{1, 2}, {4, 5}, {8, 9}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("merge(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
