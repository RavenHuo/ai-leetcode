package main

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "空树",
			root: nil,
			want: []int{},
		},
		{
			name: "单节点",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "右子树继续向下",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "完整二叉树",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
			},
			want: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name: "左斜树",
			root: &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}}},
			want: []int{5, 4, 3},
		},
		{
			name: "右斜树",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := preorderTraversal(tt.root)
			if got == nil {
				got = []int{}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
