package leetcode

/*
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ll1 := l1
	ll2 := l2
	rt := &ListNode{
		Val:  0,
		Next: nil,
	}
	var res *ListNode
	var p, a, b int
	for ll1 != nil || ll2 != nil {
		if res == nil {
			res = rt
		} else {
			res.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			res = res.Next
		}
		if ll1 != nil {
			a = ll1.Val
			ll1 = ll1.Next
		} else {
			a = 0
		}
		if ll2 != nil {
			b = ll2.Val
			ll2 = ll2.Next
		} else {
			b = 0
		}

		v := a + b + p
		if v >= 10 {
			v -= 10
			p = 1
		} else {
			p = 0
		}
		res.Val = v
	}
	res.Next = nil
	return rt
}

func AddTwoNumbersTest() {
	l1 := &ListNode{2, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{3, nil}

	l2 := &ListNode{5, nil}
	l2.Next = &ListNode{6, nil}
	l2.Next.Next = &ListNode{4, nil}

	l3 := addTwoNumbers(l1, l2)
	for ; l3 != nil; l3 = l3.Next {
		fmt.Println(l3)
	}
}
