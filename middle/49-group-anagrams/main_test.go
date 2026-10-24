package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "典型情况",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want:  [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:  "空字符串",
			input: []string{""},
			want:  [][]string{{""}},
		},
		{
			name:  "单个字符串",
			input: []string{"a"},
			want:  [][]string{{"a"}},
		},
		{
			name:  "全部属于同一组",
			input: []string{"abc", "bca", "cab"},
			want:  [][]string{{"abc", "bca", "cab"}},
		},
		{
			name:  "重复字符串",
			input: []string{"abc", "foo", "oof", "foo", "bca"},
			want:  [][]string{{"abc", "bca"}, {"foo", "oof", "foo"}},
		},
		{
			name:  "长度不同",
			input: []string{"a", "ab", "ba", "abc", "cba"},
			want:  [][]string{{"a"}, {"ab", "ba"}, {"abc", "cba"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeGroups(groupAnagrams(tt.input))
			want := normalizeGroups(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("groupAnagrams(%v) = %v, want %v", tt.input, got, want)
			}
		})
	}
}

// normalizeGroups 消除分组顺序和组内顺序不固定带来的影响。
func normalizeGroups(groups [][]string) []string {
	keys := make([]string, 0, len(groups))
	for _, group := range groups {
		words := append([]string(nil), group...)
		sort.Strings(words)

		key := ""
		for _, word := range words {
			key += word + "\x00"
		}
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return keys
}
