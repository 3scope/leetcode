package main

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	result := 0
	for len(queue) > 0 {
		// The size of each level.
		size := len(queue)
		result += size
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}

	return result
}

func countNodesRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// If root does not equals to nil, the depth of its children begin with 1.
	leftDepth, rightDepth := 1, 1
	leftNode := root.Left
	rightNode := root.Right
	for leftNode != nil {
		leftNode = leftNode.Left
		leftDepth++
	}
	for rightNode != nil {
		rightNode = rightNode.Right
		rightDepth++
	}
	if leftDepth == rightDepth {
		// 2 << 1 equals to 2 ^ 2.
		return int(math.Pow(2, float64(leftDepth))) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
