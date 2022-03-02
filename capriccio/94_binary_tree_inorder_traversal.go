package main

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	var inorder func(*TreeNode, *[]int)
	inorder = func(root *TreeNode, result *[]int) {
		if root == nil {
			return
		}
		inorder(root.Left, result)
		*result = append(*result, root.Val)
		inorder(root.Right, result)
	}

	inorder(root, &result)
	return result
}

func inorderTraversalUsingIteration(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for node := root; node != nil; node = node.Left {
		stack = append(stack, node)
	}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Right != nil {
			for node := cur.Right; node != nil; node = node.Left {
				stack = append(stack, node)
			}
		}
	}

	return result
}
