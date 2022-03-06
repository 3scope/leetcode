package main

import "math"

func isValidBST(root *TreeNode) bool {
	// Nil pointer cannot use star get the value.
	var pre = &TreeNode{
		Val: math.MinInt,
	}
	return isValidBSTRecursion(root, pre)
}

func isValidBSTRecursion(root, pre *TreeNode) bool {
	if root == nil {
		return true
	}
	left := isValidBSTRecursion(root.Left, pre)
	if pre.Val >= root.Val {
		return false
	}
	// When use slice or pointer, you should pass value through the pointer, and use star to change values of the real object.
	*pre = *root
	right := isValidBSTRecursion(root.Right, pre)

	return left && right
}

func isValidBSTIteration(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 0)
	for node := root; node != nil; node = node.Left {
		stack = append(stack, node)
	}
	var pre *TreeNode
	for len(stack) > 0 {
		size := len(stack)
		for i := 0; i < size; i++ {
			cur := stack[len(stack)-1]
			if pre != nil {
				if cur.Val <= pre.Val {
					return false
				}
			}
			stack = stack[:len(stack)-1]
			pre = cur
			if cur.Right != nil {
				for node := cur.Right; node != nil; node = node.Left {
					stack = append(stack, node)
				}
			}
		}
	}

	return true
}
