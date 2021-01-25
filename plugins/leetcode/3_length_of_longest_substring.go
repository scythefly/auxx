package main

/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	mark := make([]int, 256)
	var max int
	var cnt int
	// lastIdx 计算时的最左边
	var lastIdx int

	for idx, b := range s {
		if mark[b] > 0 {
			v := mark[b]
			if v < lastIdx {
				v = lastIdx
			}
			cnt = idx + 1 - v
			if cnt > max {
				max = cnt
			} else {
				if lastIdx < mark[b] {
					lastIdx = mark[b]
				}
			}
			mark[b] = idx + 1
			continue
		} else {
			mark[b] = idx + 1
			cnt++
			if cnt > max {
				max = cnt
			}
		}
	}
	if cnt > max {
		return cnt
	}
	return max
}

func printLengthOfLongestSubstring(s string) {
	fmt.Println(s, "=>", lengthOfLongestSubstring(s))
}

func runLengthOfLongestSubstring() {
	printLengthOfLongestSubstring("")
	printLengthOfLongestSubstring(" ")
	printLengthOfLongestSubstring("au")
	printLengthOfLongestSubstring("dvdf")
	printLengthOfLongestSubstring("cdd")
	printLengthOfLongestSubstring("abba")
	printLengthOfLongestSubstring("abcabcbb")
	printLengthOfLongestSubstring("abba")
	printLengthOfLongestSubstring("abcabcbb")
	printLengthOfLongestSubstring("bbbbb")
	printLengthOfLongestSubstring("pwwkew")
	printLengthOfLongestSubstring("wobgrovw")
	printLengthOfLongestSubstring("zwnigfunjwz")
	printLengthOfLongestSubstring("nxvloyvgmliuqandly")
}
