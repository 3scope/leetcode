package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrderN(root *Node) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*Node, 1)
	queue[0] = root
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			level[i] = cur.Val
			for j := 0; j < len(cur.Children); j++ {
				queue = append(queue, cur.Children[j])
			}
		}
		result = append(result, level)
	}

	return result
}
