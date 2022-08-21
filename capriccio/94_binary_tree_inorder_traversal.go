package main

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 1)
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		// 终止条件。
		if root == nil {
			return
		}
		// 中序遍历。
		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}
	// 执行函数。
	inorder(root)

	return result
}

func inorderTraversalIteration(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0, 1)
	// 使用栈模拟递归，中序遍历需要进行一次预处理，需要把从根节点到最左边的节点加入到栈中。
	stack := make([]*TreeNode, 0, 1)
	for node := root; node != nil; node = node.Left {
		stack = append(stack, node)
	}
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 每次进行处理的时候，都需要考虑处理节点的右孩子。
		result = append(result, cur.Val)
		if cur.Right != nil {
			for node := cur.Right; node != nil; node = node.Left {
				stack = append(stack, node)
			}
		}
	}

	return result
}
