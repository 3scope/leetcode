package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	count := 0
	return findPath(root, &count, targetSum)
}

func findPath(root *TreeNode, count *int, targetSum int) bool {
	*count += root.Val
	if root.Left == nil && root.Right == nil {
		if *count == targetSum {
			return true
		} else {
			return false
		}
	}

	left, right := false, false
	if root.Left != nil {
		left = findPath(root.Left, count, targetSum)
		*count -= root.Left.Val
	}
	if root.Right != nil {
		right = findPath(root.Right, count, targetSum)
		*count -= root.Right.Val
	}
	return left || right
}
