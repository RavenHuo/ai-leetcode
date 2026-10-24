package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "示例一", input: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{name: "示例二包含零", input: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
		{name: "两个正数", input: []int{2, 3}, want: []int{3, 2}},
		{name: "两个负数", input: []int{-1, -2}, want: []int{-2, -1}},
		{name: "两个零", input: []int{0, 0}, want: []int{0, 0}},
		{name: "正负数混合", input: []int{1, -2, 3, -4}, want: []int{24, -12, 8, -6}},
		{name: "只有一个零", input: []int{0, 2, 3}, want: []int{6, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.input...)
			if got := productExceptSelf(input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
