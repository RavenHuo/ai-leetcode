package main

/*
题目：字符串中的第一个唯一字符
LeetCode 题号：387
难度：简单
题目链接：https://leetcode.cn/problems/first-unique-character-in-a-string/description/

给定一个字符串 s，找到它的第一个不重复的字符，并返回它的索引。
如果不存在，则返回 -1。

示例 1：
输入：s = "leetcode"
输出：0

示例 2：
输入：s = "loveleetcode"
输出：2

示例 3：
输入：s = "aabb"
输出：-1

提示：
- 1 <= s.length <= 10^5
- s 只包含小写字母

函数签名：
func firstUniqChar(s string) int
*/

func firstUniqChar(s string) int {
	m := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] = append(m[s[i]], i)
	}
	for _, value := range m {
		if len(value) == 1 {
			return value[0]
		}
	}
	return -1
}
