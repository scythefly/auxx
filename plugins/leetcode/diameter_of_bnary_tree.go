package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	diameter int
)

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	diameter = 0
	treeDepth(root)
	return diameter - 1
}

func treeDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := treeDepth(node.Left)
	right := treeDepth(node.Right)
	max := left + right + 1
	if diameter < max {
		diameter = max
	}

	if left > right {
		return left + 1
	}
	return right + 1
}

/*
       1
      / \
     2   3
    / \
   4   5
  /
 6
*/
func TestDiameterOfBinaryTree() {
	t6 := &TreeNode{Val: 6}
	t4 := &TreeNode{Val: 4, Left: t6}
	t5 := &TreeNode{Val: 5}
	t3 := &TreeNode{Val: 3}

	t2 := &TreeNode{
		Val:   2,
		Left:  t4,
		Right: t5,
	}
	t1 := &TreeNode{
		Val:   1,
		Left:  t2,
		Right: t3,
	}
	fmt.Println(diameterOfBinaryTree(t1))
	r0 := unmarshalToTree([]int{1, 2, 3, 4, 5}, 10)
	fmt.Println(diameterOfBinaryTree(r0))
	//[4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]
	r2 := unmarshalToTree([]int{4, -7, -3, 10, 10, -9, -3, 9, -7, -4, 10, 6, 10, -6, -6, 10, 10, 0, 6, 5, 10, 9, 10, 10, -1, -4, 10, 10, 10, -2}, 10)
	fmt.Println(diameterOfBinaryTree(r2))
}

func unmarshalToTree(vals []int, exp int) *TreeNode {
	if len(vals) < 1 {
		return nil
	}
	var line []*TreeNode
	root := &TreeNode{Val: vals[0]}
	line = append(line, root)
	idx := 1
	for idx < len(vals) {
		line, idx = unmarshalToLine(line, idx, exp, vals)
	}

	return root
}

func unmarshalToLine(line []*TreeNode, idx, exp int, vals []int) ([]*TreeNode, int) {
	var newLine []*TreeNode
	for _, node := range line {
		if idx >= len(vals) {
			goto end
		}
		l := vals[idx]
		if l != exp {
			node.Left = &TreeNode{Val: l}
			newLine = append(newLine, node.Left)
		}
		idx++
		if idx >= len(vals) {
			goto end
		}
		r := vals[idx]
		idx++

		if r != exp {
			node.Right = &TreeNode{Val: r}
			newLine = append(newLine, node.Right)
		}
	}

end:
	return newLine, idx
}
