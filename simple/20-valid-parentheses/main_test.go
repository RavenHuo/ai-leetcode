package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "成对的小括号", input: "()", want: true},
		{name: "多种括号正确嵌套", input: "()[]{}", want: true},
		{name: "括号类型不匹配", input: "(]", want: false},
		{name: "多层嵌套", input: "([])", want: true},
		{name: "交叉嵌套", input: "([)]", want: false},
		{name: "缺少右括号", input: "((", want: false},
		{name: "多余右括号", input: "]", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.input); got != tt.want {
				t.Errorf("isValid(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
