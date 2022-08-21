package main

// 对于普通二叉树来说，使用先序 / 中序 / 后序 / 层序都可以得到所有节点。
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	result := 0
	for len(queue) > 0 {
		// The size of each level.
		size := len(queue)
		result += size
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}

	return result
}

func countNodesRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 判断它是否是一颗满二叉树，是的话直接返回总节点个数。
	leftDepth, rightDepth := 0, 0
	leftNode := root.Left
	rightNode := root.Right
	// 求出最左边的深度。
	for leftNode != nil {
		leftNode = leftNode.Left
		leftDepth++
	}
	// 求出最右边的深度。
	for rightNode != nil {
		rightNode = rightNode.Right
		rightDepth++
	}
	// 如果相同，那么它是一颗满二叉树。
	if leftDepth == rightDepth {
		return 1<<(leftDepth+1) - 1
	}

	// 先序遍历，上面处理当前节点，下面处理左右孩子。
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
