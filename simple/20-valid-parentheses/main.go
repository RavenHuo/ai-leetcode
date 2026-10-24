package main

/*
题目：有效的括号
LeetCode 题号：20
难度：简单
题目链接：https://leetcode.cn/problems/valid-parentheses/description/

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s，判断字符串是否有效。

有效字符串需满足：
1. 左括号必须用相同类型的右括号闭合。
2. 左括号必须以正确的顺序闭合。
3. 每个右括号都有一个对应的相同类型的左括号。

示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([])"
输出：true

示例 5：
输入：s = "([)]"
输出：false

提示：
1 <= s.length <= 10^4
s 仅由括号 '()[]{}' 组成。

函数签名：
func isValid(s string) bool
*/

func isValid(s string) bool {
	stack := make([]string, 0, len(s))
	for _, ch := range s {
		if string(ch) == "(" || string(ch) == "{" || string(ch) == "[" {
			stack = append(stack, string(ch))
		} else if string(ch) == ")" || string(ch) == "]" || string(ch) == "}" {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if string(ch) == ")" && last == "(" {
				stack = stack[:len(stack)-1]
			} else if string(ch) == "]" && last == "[" {
				stack = stack[:len(stack)-1]
			} else if string(ch) == "}" && last == "{" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
