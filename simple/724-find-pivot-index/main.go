package main

/*
题目：寻找数组的中心下标
LeetCode 题号：724
难度：简单
题目链接：https://leetcode.cn/problems/find-pivot-index/

给你一个整数数组 nums ，请计算数组的中心下标。

数组中心下标的定义是：其左侧所有元素相加之和等于右侧所有元素相加之和的下标。

如果中心下标位于数组最左端，那么左侧数之和视为 0；同样地，如果中心下标位于数组最右端，
那么右侧数之和视为 0。

如果数组有多个中心下标，应返回最左边的那个。如果不存在中心下标，返回 -1。

示例 1：
输入：nums = [1,7,3,6,5,6]
输出：3
解释：中心下标是 3，左侧数之和 1 + 7 + 3 = 11，右侧数之和 5 + 6 = 11。

示例 2：
输入：nums = [1,2,3]
输出：-1

示例 3：
输入：nums = [2,1,-1]
输出：0
解释：中心下标是 0，左侧数之和为 0，右侧数之和为 1 + (-1) = 0。

提示：
1 <= nums.length <= 10^4
-1000 <= nums[i] <= 1000
*/

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	leftNums := nums[left]
	rightNums := nums[right]
	for left < right {
		if leftNums >= rightNums {
			right -= 1
			rightNums += nums[right]
		} else {
			left += 1
			leftNums += nums[left]
		}
	}
	if leftNums == rightNums {
		return left
	}
	return -1
}
