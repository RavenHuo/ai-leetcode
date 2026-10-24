package main

import (
	"sort"
	"strings"
)

/*
题目：最小覆盖子串
LeetCode 题号：76
难度：困难
题目链接：https://leetcode.cn/problems/minimum-window-substring/description/

给定两个字符串 s 和 t，长度分别是 m 和 n，返回 s 中的最短窗口子串，使得该子串包含 t 中的每一个字符，包括重复字符。
如果没有这样的子串，返回空字符串 ""。

测试用例保证答案唯一。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。

示例 2：
输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。

示例 3：
输入：s = "a", t = "aa"
输出：""
解释：t 中两个字符 'a' 均应包含在 s 的子串中，因此没有符合条件的子字符串。

提示：
m == s.length
n == t.length
1 <= m, n <= 10^5
s 和 t 由英文字母组成

进阶：你能设计一个在 O(m + n) 时间内解决此问题的算法吗？

解题提示：
使用滑动窗口和字符频次统计，右指针扩展窗口直到满足覆盖条件，再移动左指针收缩窗口并更新最短答案。

函数签名：
func minWindow(s string, t string) string
*/

func minWindow(s string, t string) string {
	res := ""
	// 记录坐标
	pm := make([]string, len(s))
	// 左右指针
	left := 0
	right := left
	for right < len(s) {
		if strings.Contains(t, string(s[right])) {
			pm[right] = string(s[right])

			cur := ""
			for i := left; i <= right; i++ {
				if pm[i] != "" {
					cur += pm[i]
				}
			}
			if checkMatch(t, cur) {
				//对比当前结果
				if res == "" {
					res = s[left : right+1]
				} else if right-left < len(res) {
					res = s[left : right+1]
				}
				minP := 0
				// 找到第一个匹配的位置，并移动到此处
				for i := left + 1; i <= right; i++ {
					if pm[i] != "" {
						minP = i
						break
					}
				}
				if minP > left {
					left = minP
					continue
				}
			}

		}
		right++
	}

	return res
}

func sortString(s string) string {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}
func checkMatch(s string, t string) bool {
	sm := make(map[byte]int)
	tm := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sm[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		tm[t[i]]++
	}
	for k, v := range sm {
		tv, ok := tm[k]
		if !ok {
			return false
		}
		if tv < v {
			return false
		}
	}
	return true
}
