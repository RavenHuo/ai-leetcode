package main

/*
题目：移动零
LeetCode 题号：283
难度：简单
题目链接：https://leetcode.cn/problems/move-zeroes/description/

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意，必须在不复制数组的情况下原地对数组进行操作。

示例 1：
输入：nums = [0,1,0,3,12]
输出：[1,3,12,0,0]

示例 2：
输入：nums = [0]
输出：[0]

提示：
- 1 <= nums.length <= 10^4
- -2^31 <= nums[i] <= 2^31 - 1

进阶：你能尽量减少完成的操作次数吗？

函数签名：
func moveZeroes(nums []int)
*/

func moveZeroes(nums []int) {
	// slow 指向下一个应该放置非零元素的位置。
	slow := 0

	// fast 负责遍历数组，保证非零元素按照原来的顺序处理。
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == 0 {
			continue
		}

		if slow != fast {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		slow++
	}

}
