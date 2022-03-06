package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	result := int(math.Pow10(5))
	pre := &TreeNode{
		Val: -int(math.Pow10(5)),
	}
	getMinimumDifferenceRecursion(root, pre, &result)

	return result
}

func getMinimumDifferenceRecursion(root, pre *TreeNode, result *int) {
	if root == nil {
		return
	}
	getMinimumDifferenceRecursion(root.Left, pre, result)

	if root.Val-pre.Val < *result {
		*result = root.Val - pre.Val
	}

	*pre = *root
	getMinimumDifferenceRecursion(root.Right, pre, result)
}
