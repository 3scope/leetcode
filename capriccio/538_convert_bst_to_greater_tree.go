package main

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pre := 0
	convertBSTRecursion(root, &pre)

	return root
}

func convertBSTRecursion(root *TreeNode, pre *int) {
	if root == nil {
		return
	}
	convertBSTRecursion(root.Right, pre)
	root.Val += *pre
	*pre = root.Val
	convertBSTRecursion(root.Left, pre)
}
