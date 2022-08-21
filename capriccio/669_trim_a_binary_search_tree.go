package main

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 中序遍历。
	// 如果当前节点值在区间之外。
	if root.Val < low {
		// 当前节点的右孩子还有机会在区间中。
		return trimBST(root.Right, low, high)
	} else if root.Val > high {
		// 当前节点的左孩子还有机会在区间中。
		return trimBST(root.Left, low, high)
	}
	// 当前节点满足的话，那么修剪它的左右孩子。
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
