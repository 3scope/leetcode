package main

import "math"

// 遍历二叉树，每个节点使用“getDepth”求得左右孩子的高度。
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 1)
	stack[0] = root
	// 中序遍历。
	for node := root.Left; node != nil; node = node.Left {
		stack = append(stack, node)
	}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 得到左右孩子的高度。
		leftDepth := getDepth(cur.Left)
		rightDepth := getDepth(cur.Right)
		// 如果存在不平衡的情况，直接返回“false”。
		if math.Abs(float64(leftDepth-rightDepth)) > 1 {
			return false
		}
		for node := cur.Right; node != nil; node = node.Left {
			stack = append(stack, node)
		}
	}

	return true
}

func getDepth(cur *TreeNode) int {
	if cur == nil {
		return 0
	}

	// 后序遍历，求树的最大深度。
	left := getDepth(cur.Left)
	right := getDepth(cur.Right)
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	return max(left, right) + 1
}

func isBalancedRecursion(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if getDepthRecursion(root) != -1 {
		return true
	}
	return false
}

func getDepthRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 后序遍历，求左右孩子的深度。
	leftDepth := getDepthRecursion(root.Left)
	rightDepth := getDepthRecursion(root.Right)
	if leftDepth == -1 || rightDepth == -1 {
		return -1
	}
	if math.Abs(float64(leftDepth-rightDepth)) > 1 {
		return -1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	return max(leftDepth, rightDepth) + 1
}
