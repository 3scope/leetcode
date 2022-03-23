package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// Termination condition.
	if root1 == nil && root2 == nil {
		return nil
	}

	// Preorder traversal.
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}
