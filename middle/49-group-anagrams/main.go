package main

import "sort"

/*
题目：字母异位词分组
LeetCode 题号：49
难度：中等
题目链接：https://leetcode.cn/problems/group-anagrams/description/

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

示例 1：

输入：strs = ["eat", "tea", "tan", "ate", "nat", "bat"]

输出：[["bat"],["nat","tan"],["ate","eat","tea"]]

解释：

在 strs 中没有字符串可以通过重新排列来形成 "bat"。
字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
字符串 "ate"、"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。

示例 2：

输入：strs = [""]

输出：[[""]]

示例 3：

输入：strs = ["a"]

输出：[["a"]]

提示：

1 <= strs.length <= 10^4
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母

解题提示：
可以为每个字符串生成一个规范化的特征，例如将字符排序后作为哈希表的键；拥有相同键的字符串属于同一组。

函数签名：
func groupAnagrams(strs []string) [][]string
*/

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	res := make([][]string, 0)
	for _, str := range strs {
		sortStr := sortString(str)
		if _, ok := m[sortStr]; !ok {
			m[sortStr] = make([]string, 0)
		}
		m[sortStr] = append(m[sortStr], str)
	}
	for _, strArr := range m {
		res = append(res, strArr)
	}
	return res
}

func sortString(str string) string {
	b := []byte(str)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
