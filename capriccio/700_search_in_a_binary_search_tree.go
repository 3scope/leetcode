package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

func searchBSTIteration(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	for root != nil {
		if root.Val < val {
			root = root.Left
		} else if root.Val == val {
			return root
		} else {
			root = root.Right
		}
	}
	return nil
}
