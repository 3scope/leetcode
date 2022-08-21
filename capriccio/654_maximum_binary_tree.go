package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 求出最大值，根据最大值得到对应的节点。
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[maxIndex] < nums[i] {
			maxIndex = i
		}
	}

	leftNums := nums[:maxIndex]
	rightNums := nums[maxIndex+1:]
	root := new(TreeNode)
	root.Val = nums[maxIndex]
	root.Left = constructMaximumBinaryTree(leftNums)
	root.Right = constructMaximumBinaryTree(rightNums)
	return root
}
