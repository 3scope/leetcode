package main

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// Middle.
	if root.Val < low {
		return trimBST(root.Right, low, high)
	} else if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	// Left.
	root.Left = trimBST(root.Left, low, high)
	// Rigth.
	root.Right = trimBST(root.Right, low, high)

	return root
}
