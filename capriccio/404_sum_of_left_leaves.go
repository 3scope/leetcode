package main

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		result += root.Left.Val
	}

	left := sumOfLeftLeaves(root.Left)
	right := sumOfLeftLeaves(root.Right)
	result += left + right

	return result
}
