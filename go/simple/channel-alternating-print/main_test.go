package main

import "testing"

func TestAlternatePrint(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{name: "零次输出", input: 0, want: ""},
		{name: "负数输入", input: -1, want: ""},
		{name: "输出一组", input: 1, want: "A1"},
		{name: "输出三组", input: 3, want: "A1B2C3"},
		{name: "输出十组", input: 10, want: "A1B2C3D4E5F6G7H8I9J10"},
		{name: "输出全部字母", input: 26, want: "A1B2C3D4E5F6G7H8I9J10K11L12M13N14O15P16Q17R18S19T20U21V22W23X24Y25Z26"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternatePrint(tt.input); got != tt.want {
				t.Errorf("alternatePrint(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
