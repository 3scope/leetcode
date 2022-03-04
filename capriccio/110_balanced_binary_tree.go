package main

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 1)
	stack[0] = root
	for node := root.Left; node != nil; node = node.Left {
		stack = append(stack, node)
	}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		leftDepth := getDepth(cur.Left)
		rightDepth := getDepth(cur.Right)
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
	queue := make([]*TreeNode, 1)
	queue[0] = cur
	depth := 0
	for len(queue) > 0 {
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
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
	leftDepth := getDepthRecursion(root.Left)
	rightDepth := getDepthRecursion(root.Right)
	if leftDepth == -1 || rightDepth == -1 {
		return -1
	}
	if math.Abs(float64(leftDepth-rightDepth)) > 1 {
		return -1
	}
	return 1 + int(math.Max(float64(leftDepth), float64(rightDepth)))

}
