package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) > 0 {
		// The size of each level.
		size := len(queue)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			level[i] = cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
