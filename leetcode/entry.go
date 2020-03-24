package leetcode

import "fmt"

var (
	latest int = 10
)

const (
	INT32_MIN = -(2 << 30)
	INT32_MAX = 2<<30 - 1
)

func Do(idx int) {
	if idx == 0 {
		idx = latest
	}
	switch idx {
	case 1:
		TestTwoSum()
	case 2:
		AddTwoNumbersTest()
	case 3:
		TestLengthOfLongestSubstring()
	case 4:
		FindMedianSortedArraysTest()
	case 5:
		TestLongestPalindrome()
	case 6:
		TestConvert()
	case 7:
		TestReverse()
	case 8:
		TestMyAtoi()
	case 9:
		TestIsPalindrome()
	case 10:
		TestIsMatch()
	case 1286:
		TestCombinationIterator()
	default:
		fmt.Println("To Be Continued...")
	}
	// LargestComponentSizeTest()
	//TestDiameterOfBinaryTree()
}
