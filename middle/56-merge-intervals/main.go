package main

import "sort"

/*
题目：合并区间
LeetCode 题号：56
难度：中等
题目链接：https://leetcode.cn/problems/merge-intervals/

以数组 intervals 表示若干个区间集合，其中单个区间为 intervals[i] = [starti, endi]。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]

提示：
1 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^4

解题提示：
先按照区间左端点排序。依次遍历时，如果当前区间与结果中的最后一个区间重叠，
就更新最后一个区间的右端点；否则将当前区间加入结果。
*/

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return res
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		arr := intervals[i]
		lastRes := res[len(res)-1]
		if arr[0] <= lastRes[len(lastRes)-1] && arr[1] >= lastRes[0] {
			res[len(res)-1] = mergeArr(res[len(res)-1], arr)
		} else {
			res = append(res, arr)
		}
	}
	return res
}

func mergeArr(a, b []int) []int {
	al, ar := a[0], a[len(a)-1]
	bl, br := b[0], b[len(a)-1]
	l := min(al, bl)
	r := max(ar, br)
	return []int{l, r}
}
