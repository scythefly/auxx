package leetcode

import "fmt"

/*
7. 整数反转
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21

注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverse(x int) int {
	if x < INT32_MIN || x > INT32_MAX {
		return 0
	}
	var signed int = 1
	if x < 0 {
		signed = -1
		x = -x
	}
	if x < 10 {
		return x
	}
	var n int = 1
	var ns []int
	for {
		if x/n < 1 {
			break
		}
		ns = append(ns, x%(n*10)/n)
		n *= 10
	}
	var ret int
	for i := 0; i < len(ns); i++ {
		ret *= 10
		ret += ns[i]
	}
	ret *= signed
	if ret < INT32_MIN || ret > INT32_MAX {
		return 0
	}
	return ret
}

func printReverse(x int) {
	fmt.Println(x, "=>", reverse(x))
}

func TestReverse() {
	printReverse(123)
	printReverse(-123)
	printReverse(230)
	printReverse(-230)
	printReverse(1534236469)
	printReverse(1563847412)
}
