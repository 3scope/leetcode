package main

func invertTree(root *TreeNode) *TreeNode {
	// Do nothing.
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func invertTreeIteration(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 1)
	stack[0] = root
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		cur.Left, cur.Right = cur.Right, cur.Left
	}
}
