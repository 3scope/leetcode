package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 终止条件。
	if root1 == nil && root2 == nil {
		return nil
	}

	// 将“root2”合并到“root1”上。
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 如果都存在，就将相应的值加过去。
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}
