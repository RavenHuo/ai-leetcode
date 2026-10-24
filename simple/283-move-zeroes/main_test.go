package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "示例一", input: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		{name: "示例二", input: []int{0}, want: []int{0}},
		{name: "全部为零", input: []int{0, 0, 0}, want: []int{0, 0, 0}},
		{name: "没有零", input: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "零在末尾", input: []int{1, 2, 0}, want: []int{1, 2, 0}},
		{name: "负数和零", input: []int{0, -1, 0, -2}, want: []int{-1, -2, 0, 0}},
		{name: "零在开头", input: []int{0, 0, 1, 2}, want: []int{1, 2, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.input...)
			moveZeroes(input)
			if !reflect.DeepEqual(input, tt.want) {
				t.Errorf("moveZeroes(%v) = %v, want %v", tt.input, input, tt.want)
			}
		})
	}
}
