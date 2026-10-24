package main

/*
题目：连续的子数组和
LeetCode 题号：523
难度：中等
题目链接：https://leetcode.cn/problems/continuous-subarray-sum/description/

给你一个整数数组 nums 和一个整数 k ，如果 nums 有一个 好的子数组 返回 true ，否则返回 false：

一个 好的子数组 是：

    长度至少为 2 ，且
    子数组元素总和为 k 的倍数。

注意：

    子数组是数组中连续的部分。
    如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。
    0 始终视为 k 的一个倍数。

示例 1：

输入：nums = [23,2,4,6,7], k = 6
输出：true
解释：[2,4] 是一个大小为 2 的子数组，并且和为 6 。

示例 2：

输入：nums = [23,2,6,4,7], k = 6
输出：true
解释：[23, 2, 6, 4, 7] 是大小为 5 的子数组，并且和为 42 。
42 是 6 的倍数，因为 42 = 7 * 6 且 7 是一个整数。

示例 3：

输入：nums = [23,2,6,4,7], k = 13
输出：false

提示：

    1 <= nums.length <= 10^5
    0 <= nums[i] <= 10^9
    0 <= sum(nums[i]) <= 2^31 - 1
    1 <= k <= 2^31 - 1

解题提示：
可以观察连续子数组的前缀和对 k 取模后的结果；如果两个不同位置的前缀和模 k 相同，
它们之间的子数组和就是 k 的倍数。需要特别保证子数组长度至少为 2，并保留每个余数第一次出现的位置。

函数签名：
func checkSubarraySum(nums []int, k int) bool
*/

func checkSubarraySum(nums []int, k int) bool {
	// 记录前缀和 mod k后的余数值, value 是对应的下标值
	m := make(map[int][]int)
	prefix := 0
	for i, num := range nums {
		prefix += num
		mod := prefix % k
		if len(m[mod]) == 0 {
			m[mod] = make([]int, 0)
			m[mod] = append(m[mod], i)
		} else {
			// 长度至少为 2
			if i-m[mod][0] >= 1 {
				return true
			}
			m[mod] = append(m[mod], i)
		}
	}
	return false
}
