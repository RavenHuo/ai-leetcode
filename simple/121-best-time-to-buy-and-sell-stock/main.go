package main

/*
题目：买卖股票的最佳时机
LeetCode 题号：121
难度：简单
题目链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

给定一个数组 prices，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择某一天买入这只股票，并选择在未来的某一个不同的日子卖出该股票。
设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0。

示例 1：
输入：prices = [7,1,5,3,6,4]
输出：5
解释：在第 2 天买入，在第 5 天卖出，最大利润为 6 - 1 = 5。

示例 2：
输入：prices = [7,6,4,3,1]
输出：0
解释：没有交易完成，所以最大利润为 0。

示例 3：
输入：prices = [2,4,1]
输出：2
解释：在价格为 2 时买入，在价格为 4 时卖出，最大利润为 2。

提示：
1 <= prices.length <= 10^5
0 <= prices[i] <= 10^4

函数签名：
func maxProfit(prices []int) int
*/

func maxProfit(prices []int) int {
	maxP := 0
	minPrice := prices[0]
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		}
		curP := price - minPrice
		if curP > maxP {
			maxP = curP
		}
	}
	return maxP
}
