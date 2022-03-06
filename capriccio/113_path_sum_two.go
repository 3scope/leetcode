package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	path := make([]int, 1)
	path[0] = root.Val
	targetSum = targetSum - root.Val
	getPathSum(root, targetSum, &path, &result)
	return result
}

func getPathSum(root *TreeNode, targetSum int, path *[]int, result *[][]int) {
	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			temp := make([]int, len(*path))
			copy(temp, *path)
			*result = append(*result, temp)
		}
		return
	}
	if root.Left != nil {
		targetSum -= root.Left.Val
		*path = append(*path, root.Left.Val)
		getPathSum(root.Left, targetSum, path, result)
		*path = (*path)[0 : len(*path)-1]
		targetSum += root.Left.Val
	}
	if root.Right != nil {
		targetSum -= root.Right.Val
		*path = append(*path, root.Right.Val)
		getPathSum(root.Right, targetSum, path, result)
		*path = (*path)[0 : len(*path)-1]
		targetSum += root.Right.Val
	}
}
