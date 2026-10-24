package main

/*
题目：判断是不是平衡二叉树
牛客题目标识：BM36
难度：中等
题目链接：https://www.nowcoder.com/practice/8b3b95850edb4115918ecebdf1b4d222

输入一棵节点数为 n 的二叉树，判断该二叉树是否是平衡二叉树。
在这里，只需要考虑其平衡性，不需要考虑它是不是排序二叉树。

平衡二叉树具有以下性质：它是一棵空树，或者它的左右两个子树的高度差的绝对值不超过 1，
并且左右两个子树都是平衡二叉树。空树约定为平衡二叉树。

输入格式：
输入一棵二叉树的根节点 pRoot。

输出格式：
输出一个布尔类型的值，表示该二叉树是否为平衡二叉树。

示例 1：
输入：{1,2,3,4,5,6,7}
输出：true

示例 2：
输入：{}
输出：true

示例 3：
输入：{1,#,2,#,3}
输出：false

数据范围：
- 0 <= n <= 1000
- 树上节点的 val 值为整数。

解题提示：
可以在计算子树高度的同时判断是否平衡。若某棵子树不平衡，就返回一个特殊高度，
让不平衡状态向上传递，避免重复计算每个节点的高度。

函数签名：
func isBalanced(root *TreeNode) bool
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := hightTree(root.Left)
	right := hightTree(root.Right)
	if !(left-right <= 1 && left-right >= -1) {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func hightTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := hightTree(root.Left)
	right := hightTree(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
