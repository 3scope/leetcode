package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[maxIndex] < nums[i] {
			maxIndex = i
		}
	}
	// Preorder traversal.
	leftNums := nums[:maxIndex]
	rightNums := nums[maxIndex+1:]
	root := new(TreeNode)
	root.Val = nums[maxIndex]
	root.Left = constructMaximumBinaryTree(leftNums)
	root.Right = constructMaximumBinaryTree(rightNums)
	return root
}
