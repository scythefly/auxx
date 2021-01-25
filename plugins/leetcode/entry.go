package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	latest int = 10
)

const (
	INT32_MIN = -(2 << 30)
	INT32_MAX = 2<<30 - 1
)

func run(idx int) {
	if idx == 0 {
		idx = latest
	}
	switch idx {
	case 1:
		runTwoSum()
	case 2:
		runAddTwoNumbersTest()
	case 3:
		runLengthOfLongestSubstring()
	case 4:
		runFindMedianSortedArraysTest()
	case 5:
		runLongestPalindrome()
	case 6:
		runConvert()
	case 7:
		runReverse()
	case 8:
		runMyAtoi()
	case 9:
		runIsPalindrome()
	case 10:
		runIsMatch()
	case 1286:
		runCombinationIterator()
	default:
		fmt.Println("To Be Continued...")
	}
	// LargestComponentSizeTest()
	// TestDiameterOfBinaryTree()
}

// func newLeetcodeCommand() cobra.Command {
// 	cmd := cobra.Command{
// 		Use:   "leetcode",
// 		Short: "Do leetcode",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			var index int
// 			if len(args) > 0 {
// 				index, _ = strconv.Atoi(args[0])
// 			}
// 			run(index)
// 		},
// 	}
//
// 	return cmd
// }

// Var Exported in Plugin
// var Commander = newLeetcodeCommand()

var Commander = cobra.Command{
	Use:   "leetcode",
	Short: "Do leetcode",
	Run: func(cmd *cobra.Command, args []string) {
		var index int
		if len(args) > 0 {
			index, _ = strconv.Atoi(args[0])
		}
		run(index)
	},
}
