package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := new(TreeNode)
	index := len(nums) / 2
	root.Val = nums[index]

	root.Left = sortedArrayToBST(nums[:index])
	root.Right = sortedArrayToBST(nums[index+1:])

	return root
}
