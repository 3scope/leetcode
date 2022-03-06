package main

func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*TreeNode, 1)
	queue[0] = root
	result := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if i == 0 {
				result = cur.Val
			}
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
