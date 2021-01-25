package main

import "fmt"

/*
4. 寻找两个有序数组的中位数
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	mid := total / 2

	// make sure len(nums1) > len(nums2)
	if len(nums1) < len(nums2) {
		t := nums1
		nums1 = nums2
		nums2 = t
	}

	if len(nums2) == 0 {
		if total%2 == 0 {
			return float64(nums1[mid]+nums1[mid-1]) / 2.0
		}
		return float64(nums1[mid])
	}

	var i1, i2, it int
	var ii1, ii2 int

	var c1, ct int
	it = -1
	if total%2 == 0 {
		for {
			for i := ii1; i < len(nums1); i++ {
				i1 = i
				if nums1[i] <= nums2[i2] || ii2 == len(nums2) {
					it++
					ii1++
					if ct == 1 {
						return float64(c1+nums1[i]) / 2.0
					}
					if it == mid-1 {
						c1 = nums1[i]
						ct = 1
					}
				} else {
					break
				}
			}
			for i := ii2; i < len(nums2); i++ {
				i2 = i
				if nums2[i] < nums1[i1] || ii1 == len(nums1) {
					it++
					ii2++
					if ct == 1 {
						return float64(c1+nums2[i]) / 2.0
					}
					if it == mid-1 {
						c1 = nums2[i]
						ct = 1
					}
				} else {
					break
				}
			}
		}
	}

	for {
		for i := ii1; i < len(nums1); i++ {
			i1 = i
			if nums1[i] <= nums2[i2] || ii2 == len(nums2) {
				it++
				ii1++
				if it == mid {
					return float64(nums1[i])
				}
			} else {
				break
			}
		}
		for i := ii2; i < len(nums2); i++ {
			i2 = i
			if nums2[i] < nums1[i1] || ii1 == len(nums1) {
				it++
				ii2++
				if it == mid {
					return float64(nums2[i])
				}
			} else {
				break
			}
		}
	}
}

func runFindMedianSortedArraysTest() {
	n1 := []int{1, 3, 5, 7}
	n2 := []int{2, 4, 6}

	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1, 2, 3, 4, 5}
	n2 = []int{6, 7}
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1, 2, 6, 7}
	n2 = []int{3, 4, 5}
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1, 3, 5, 7, 8}
	n2 = []int{2, 4, 6}

	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1, 2, 3, 4, 5, 8}
	n2 = []int{6, 7}
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1, 2, 6, 7, 8}
	n2 = []int{3, 4, 5}
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1, 2}
	n2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{0, 0}
	n2 = []int{0, 0}
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1}
	n2 = []int{2, 3}
	fmt.Println(findMedianSortedArrays(n1, n2))

	n1 = []int{1}
	n2 = []int{2, 3, 4}
	fmt.Println(findMedianSortedArrays(n1, n2))
}
