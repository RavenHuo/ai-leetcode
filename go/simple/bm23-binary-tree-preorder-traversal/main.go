package main

/*
题目：二叉树的前序遍历
牛客题目标识：BM23
难度：简单
题目链接：https://www.nowcoder.com/practice/5e2135f4d2b14eb8a5b06fab4c938635

给你二叉树的根节点 root，返回它节点值的前序遍历。

前序遍历的访问顺序是：根节点、左子树、右子树。

示例 1：
输入：{1,#,2,3}
输出：[1,2,3]

示例 2：
输入：{1,2,3,4,5}
输出：[1,2,4,5,3]

示例 3：
输入：{}
输出：[]

提示：
二叉树节点数量满足题目要求。
二叉树节点值为整数，且各节点值不相同。

函数签名：
func preorderTraversal(root *TreeNode) []int
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	res = append(res, root.Val)
	if root.Left != nil {
		left := root.Left
		res = append(res, preorderTraversal(left)...)
	}
	if root.Right != nil {
		right := root.Right
		res = append(res, preorderTraversal(right)...)
	}
	return res
}
