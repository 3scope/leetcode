package main

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 1)
	// 定义函数。
	var postorder func(*TreeNode)
	postorder = func(root *TreeNode) {
		// 终止条件。
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		// 后序遍历，需要最后对节点进行处理。
		result = append(result, root.Val)
	}
	postorder(root)
	return result
}

func postorderTraversalIteration(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0, 1)
	// 使用栈来模拟递归。
	stack := make([]*TreeNode, 1)
	stack[0] = root
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 对当前节点进行处理。
		result = append(result, cur.Val)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}
	// Result 中存放的结果是“中右左”的形式，因此想要获得“左右中”的后序遍历的结果，需要翻转一次数组。
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
