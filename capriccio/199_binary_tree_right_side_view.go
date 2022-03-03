package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	total := make([][]int, 0)
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) > 0 {
		level := make([]int, 0)
		// The size of each level.
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		total = append(total, level)
	}

	result := make([]int, 0)
	for i := 0; i < len(total); i++ {
		result = append(result, total[i][len(total[i])-1])
	}
	return result
}
