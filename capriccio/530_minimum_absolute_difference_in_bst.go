package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	result := int(math.Pow10(5))
	// “pre”设置为一个小数，作为初始值。
	var pre int = -1e5
	getMinimumDifferenceRecursion(root, &pre, &result)

	return result
}

func getMinimumDifferenceRecursion(root *TreeNode, pre *int, result *int) {
	if root == nil {
		return
	}
	getMinimumDifferenceRecursion(root.Left, pre, result)

	if root.Val-(*pre) < *result {
		*result = root.Val - (*pre)
	}

	*pre = root.Val
	getMinimumDifferenceRecursion(root.Right, pre, result)
}
