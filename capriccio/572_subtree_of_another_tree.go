package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	sameNode := findNode(root, subRoot)
	for i := 0; i < len(sameNode); i++ {
		if compareNode(sameNode[i], subRoot) {
			return true
		}
	}
	return false
}

func findNode(root, subRoot *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	// Preorder traversal.
	subResult := make([]*TreeNode, 0)
	if root.Val == subRoot.Val {
		subResult = append(subResult, root)
	}
	left := findNode(root.Left, subRoot)
	if left != nil {
		subResult = append(subResult, left...)
	}
	right := findNode(root.Right, subRoot)
	if right != nil {
		subResult = append(subResult, right...)
	}

	return subResult
}

func compareNode(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	// Preorder traversal.
	if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	left := compareNode(p.Left, q.Left)
	right := compareNode(p.Right, q.Right)

	return left && right
}
