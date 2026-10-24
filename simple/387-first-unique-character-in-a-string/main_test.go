package main

import "testing"

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "示例一", input: "leetcode", want: 0},
		{name: "示例二", input: "loveleetcode", want: 2},
		{name: "没有唯一字符", input: "aabb", want: -1},
		{name: "唯一字符在末尾", input: "aabbc", want: 4},
		{name: "唯一字符在中间", input: "aabbccdef", want: 6},
		{name: "只有一个字符", input: "z", want: 0},
		{name: "多个唯一字符取第一个", input: "abcd", want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.input); got != tt.want {
				t.Errorf("firstUniqChar(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
