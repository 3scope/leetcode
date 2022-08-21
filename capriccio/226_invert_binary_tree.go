package main

func invertTree(root *TreeNode) *TreeNode {
	// 特殊情况。
	if root == nil {
		return nil
	}

	preorderInvert(root)
	return root
}

func preorderInvert(root *TreeNode) {
	// 递归终止条件。
	if root == nil {
		return
	}
	// 先序遍历。
	root.Left, root.Right = root.Right, root.Left
	// 虽然左右孩子更换了位置，但是先序遍历还没有遍历孩子，所以不会造成重复。
	preorderInvert(root.Left)
	preorderInvert(root.Right)
}
