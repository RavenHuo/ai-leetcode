package main

/*
题目：最大子数组和
LeetCode 题号：53
难度：中等
题目链接：https://leetcode.cn/problems/maximum-subarray/description/

给你一个整数数组 nums，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组是数组中的一个连续部分。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6。

示例 2：
输入：nums = [1]
输出：1

示例 3：
输入：nums = [5,4,-1,7,8]
输出：23

提示：
- 1 <= nums.length <= 10^5
- -10^4 <= nums[i] <= 10^4

进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

解题提示：
遍历数组时维护“以当前位置结尾的最大子数组和”。如果之前的累计和小于 0，就从当前元素重新开始；同时记录遍历过程中出现的最大值。

函数签名：
func maxSubArray(nums []int) int
*/

func maxSubArray(nums []int) int {
	maxSub := make([]int, len(nums))
	m := nums[0]
	maxSub[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if maxSub[i-1] > 0 {
			maxSub[i] = maxSub[i-1] + nums[i]
		} else {
			maxSub[i] = nums[i]
		}
		if maxSub[i] > m {
			m = maxSub[i]
		}
	}
	return m
}
