package main

func maxDepth(root *TreeNode) int {
	result := 0
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) > 0 {
		result++
		size := len(queue)
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
}
