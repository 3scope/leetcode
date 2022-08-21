package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	count := root.Val
	return findPath(root, count, targetSum)
}

func findPath(root *TreeNode, count int, targetSum int) bool {
	// 递归终止条件，收集结果。
	if root.Left == nil && root.Right == nil {
		if count == targetSum {
			return true
		} else {
			return false
		}
	}

	left, right := false, false
	if root.Left != nil {
		left = findPath(root.Left, count+root.Left.Val, targetSum)
	}
	// 如果存在，可以直接返回。
	if left {
		return true
	}

	if root.Right != nil {
		right = findPath(root.Right, count+root.Right.Val, targetSum)
	}
	if right {
		return true
	}

	return left || right
}
