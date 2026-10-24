package main

/*
题目：除了自身以外数组的乘积
LeetCode 题号：238
难度：中等
题目链接：https://leetcode.cn/problems/product-of-array-except-self/description/

给你一个整数数组 nums，返回数组 answer，其中 answer[i] 等于 nums 中除了 nums[i] 之外其余各元素的乘积。

题目数据保证数组 nums 中任意元素的全部前缀元素和后缀的乘积都在 32 位整数范围内。
请不要使用除法，且在 O(n) 时间复杂度内完成此题。

示例 1：
输入：nums = [1,2,3,4]
输出：[24,12,8,6]

示例 2：
输入：nums = [-1,1,0,-3,3]
输出：[0,0,9,0,0]

提示：
- 2 <= nums.length <= 10^5
- -30 <= nums[i] <= 30
- 输入保证数组 answer[i] 在 32 位整数范围内。

进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？
（出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

解题提示：
先考虑每个位置左侧所有元素的乘积和右侧所有元素的乘积，再想办法用输出数组保存其中一部分结果，避免使用除法。

函数签名：
func productExceptSelf(nums []int) []int
*/

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0] = nums[0]
	right[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i] * left[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = nums[i] * right[i+1]
	}
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res[i] = right[i+1]
		} else if i == len(nums)-1 {
			res[i] = left[i-1]
		} else {
			res[i] = left[i-1] * right[i+1]
		}
	}
	return res
}
