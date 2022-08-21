package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	depth := 0
	for len(queue) > 0 {
		size := len(queue)
		depth++
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			// 当出现叶子节点，证明已经到了最小深度，可以直接返回。
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
		}
	}
	// 所有的情况都不会使用这个 Return。
	return depth
}

// 使用递归的话，逻辑稍微复杂一些，需要判断当前节点是否是叶子节点。
func minDepthRecursion(root *TreeNode) int {
	// 空节点的高度为零。
	if root == nil {
		return 0
	}

	// 后序遍历。
	left := minDepthRecursion(root.Left)
	right := minDepthRecursion(root.Right)
	// 如果只有一边为空，这个节点不是最低点，最低点需要是叶子节点。
	if left == 0 {
		return right + 1
	} else if right == 0 {
		return left + 1
	}

	// 这种情况下，只有两个孩子节点都不为空。
	if left > right {
		return right + 1
	} else {
		return left + 1
	}
}
