package main

func maxDepthNAry(root *Node) int {
	if root == nil {
		return 0
	}
	queue := make([]*Node, 1)
	queue[0] = root
	depth := 0
	for len(queue) > 0 {
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if len(cur.Children) > 0 {
				queue = append(queue, cur.Children...)
			}
		}
	}
	return depth
}
