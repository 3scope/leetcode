package main

import "math"

func findMode(root *TreeNode) []int {
	result := make([]int, 0)
	pre := &TreeNode{
		Val: -int(math.Pow10(5)) - 1,
	}
	count := 0
	maxCount := 1
	findModeInorder(root, pre, &result, &count, &maxCount)

	return result
}

func findModeInorder(root, pre *TreeNode, result *[]int, count, maxCount *int) {
	if root == nil {
		return
	}
	findModeInorder(root.Left, pre, result, count, maxCount)
	if pre.Val == -int(math.Pow10(5))-1 {
		*count = 1
	} else if root.Val == pre.Val {
		(*count)++
	} else {
		*count = 1
	}
	*pre = *root
	if *count == *maxCount {
		*result = append(*result, root.Val)
	}
	if *count > *maxCount {
		*maxCount = *count
		*result = (*result)[0:1]
		(*result)[0] = root.Val
	}
	findModeInorder(root.Right, pre, result, count, maxCount)
}
