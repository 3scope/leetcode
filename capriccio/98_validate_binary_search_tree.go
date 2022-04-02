package main

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	pre := -int(math.Pow(2, 32))
	return isValidBSTRecursion(root, &pre)
}

func isValidBSTRecursion(root *TreeNode, pre *int) bool {
	if root == nil {
		return true
	}

	left := isValidBSTRecursion(root.Left, pre)

	if *pre >= root.Val {
		return false
	}
	*pre = root.Val

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

	return true
}
