package main

// 左叶子不同于左视图，一层可以有多个左叶子。
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	// 父节点可以判断左孩子是不是左叶子，自身由于不知道自己是左孩子还是右孩子，因此无法求左叶子。
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		result += root.Left.Val
	}

	// 求子树的左叶子节点。
	left := sumOfLeftLeaves(root.Left)
	right := sumOfLeftLeaves(root.Right)
	result += left + right

	return result
}
