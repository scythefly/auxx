package leetcode

import "fmt"

/*
9. 回文数
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var bits int = 1
	var n int = 1
	for ; bits < 11; bits++ {
		n *= 10
		if x/n < 1 {
			break
		}
	}
	var idx = 0
	ct := []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}
	for {
		rx := bits - 1 - idx
		if idx >= rx {
			return true
		}
		r := x % (ct[idx] * 10) / ct[idx]
		l := x % (ct[rx] * 10) / ct[rx]
		if l != r {
			return false
		}
		idx++
	}
}

func printIsPalindrome(x int) {
	fmt.Println(x, "=>", isPalindrome(x))
}

func TestIsPalindrome() {
	printIsPalindrome(121)
	printIsPalindrome(1212)
	printIsPalindrome(1221)
	printIsPalindrome(-121)
	printIsPalindrome(12345)
	printIsPalindrome(12321)
}
