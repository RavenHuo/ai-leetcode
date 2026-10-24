package main

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "典型情况",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "整个字符串就是答案",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "字符数量不足",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "重复字符",
			s:    "aa",
			t:    "aa",
			want: "aa",
		},
		{
			name: "答案位于末尾",
			s:    "ab",
			t:    "b",
			want: "b",
		},
		{
			name: "窗口需要保留重复字符",
			s:    "bba",
			t:    "ab",
			want: "ba",
		},
		{
			name: "无匹配字符",
			s:    "abc",
			t:    "z",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.s, tt.t); got != tt.want {
				t.Errorf("minWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
