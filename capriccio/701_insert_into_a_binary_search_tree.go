package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = new(TreeNode)
		root.Val = val
		return root
	}

	if val >= root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}

func insertIntoBSTIteration(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = new(TreeNode)
		root.Val = val
		return root
	}
	pre, cur := (*TreeNode)(nil), root
	for cur != nil {
		pre = cur
		if val >= cur.Val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	cur = new(TreeNode)
	cur.Val = val

	if val > pre.Val {
		pre.Right = cur
	} else {
		pre.Left = cur
	}

	return root
}
