package main

/*
题目：和为 K 的子数组
LeetCode 题号：560
难度：中等
题目链接：https://leetcode.cn/problems/subarray-sum-equals-k/

给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。

子数组是数组中元素的连续非空序列。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2

示例 2：
输入：nums = [1,2,3], k = 3
输出：2

提示：
1 <= nums.length <= 2 * 10^4
-1000 <= nums[i] <= 1000
-10^7 <= k <= 10^7

解题提示：
可以利用前缀和。若当前前缀和为 sum，那么之前出现过的 sum-k
对应的前缀和位置，都能与当前位置组成一个和为 k 的连续子数组。
使用 map 记录每种前缀和出现的次数。
*/

func subarraySum(nums []int, k int) int {
	prefixSum := 0
	result := 0
	prefixCount := map[int]int{0: 1}

	for _, num := range nums {
		prefixSum += num
		result += prefixCount[prefixSum-k]
		prefixCount[prefixSum]++
	}

	return result
}
