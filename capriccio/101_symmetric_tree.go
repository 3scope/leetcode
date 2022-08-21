package main

func isSymmetric(root *TreeNode) bool {
	// 特殊情况。
	if root == nil {
		return true
	}

	return compareTreeNode(root.Left, root.Right)
}

func compareTreeNode(p, q *TreeNode) bool {
	// 终止条件。
	if p == nil && q == nil {
		return true
	}

	// 先序遍历，先处理当前节点。
	if p == nil {
		// P 为空，Q 不为空。
		return false
	} else if q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	// 如果出现不匹配，可以提前终止，这样的话无需遍历整个树。
	if !compareTreeNode(p.Left, q.Right) {
		return false
	}
	if !compareTreeNode(p.Right, q.Left) {
		return false
	}

	return true
}
