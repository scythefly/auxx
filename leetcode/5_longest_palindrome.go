package leetcode

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var length1, length2, idx1, idx2 int
	for i := 1; i < len(s); i++ {
		l := getPalindromeLength1(s, i)
		if l > length1 {
			length1 = l
			idx1 = i
		}
	}

	for i := 0; i < len(s); i++ {
		l := getPalindromeLength2(s, i)
		if l > length2 {
			length2 = l
			idx2 = i
		}
	}

	if length1 > length2 {
		p := idx1 - (length1-1)/2
		return s[p : p+length1]
	}
	p := idx2 - (length2-1)/2
	return s[p : p+length2]
}

func getPalindromeLength1(s string, idx int) int {
	var length int = 1
	for i := 1; idx-i >= 0 && idx+i < len(s); i++ {
		if s[idx-i] == s[idx+i] {
			length += 2
		} else {
			break
		}
	}
	return length
}

func getPalindromeLength2(s string, idx int) int {
	var length int
	for i := 0; idx-i >= 0 && idx+i+1 < len(s); i++ {
		if s[idx-i] == s[idx+i+1] {
			length += 2
		} else {
			break
		}
	}
	return length
}

func printLongestPalindrome(s string) {
	fmt.Println(s, "=>", longestPalindrome(s))
}

func TestLongestPalindrome() {
	printLongestPalindrome("abab")
	printLongestPalindrome("cbbc")
	printLongestPalindrome("fcbbc")
	printLongestPalindrome("fcbabcd")
}
