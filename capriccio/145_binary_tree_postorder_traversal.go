package main

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	var postorder func(*TreeNode, *[]int)
	postorder = func(root *TreeNode, result *[]int) {
		if root == nil {
			return
		}
		postorder(root.Left, result)
		postorder(root.Right, result)
		*result = append(*result, root.Val)
	}
	postorder(root, &result)

	return result
}
