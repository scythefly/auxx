package leetcode

import "fmt"

var (
	latest int = 6
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
	default:
		fmt.Println("To Be Continue...")
	}
	// LargestComponentSizeTest()
	//TestDiameterOfBinaryTree()
}
