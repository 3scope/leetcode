package main

func levelOrder(root *TreeNode) [][]int {
	// 特殊情况。
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	// 使用队列保存每个节点，保证每层从左到右的顺序。
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) > 0 {
		size := len(queue)
		// 新申请一个Slice，用于保存每层的数据。
		subResult := make([]int, 0, 1)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			subResult = append(subResult, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 保存每层的数据。
		result = append(result, subResult)
	}

	return result
}
