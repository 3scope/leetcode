package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// 特殊情况。
	if subRoot == nil {
		return true
	}
	// 终止条件。
	if root == nil {
		return false
	}

	// 先序遍历，首先判断当前节点作为根节点，是否能和 SubRoot 相等。
	// 先判断以当前节点为根节点，之后判断左右孩子。
	if compareSameTree(root, subRoot) {
		return true
	}
	// 只要左右孩子中有一个满足条件，就证明子树中含有 SubRoot.
	if isSubtree(root.Left, subRoot) {
		return true
	}
	if isSubtree(root.Right, subRoot) {
		return true
	}

	return false
}

func compareSameTree(p, q *TreeNode) bool {
	// 终止条件。
	if p == nil && q == nil {
		return true
	}

	// 先序遍历，判断当前的连个节点。
	if p == nil {
		return false
	} else if q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	// 比较子节点，当不相同时可以直接返回。
	if !compareSameTree(p.Left, q.Left) {
		return false
	}
	if !compareSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
