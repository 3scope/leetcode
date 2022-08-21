package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用层序遍历找到树的左下角的值。
func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*TreeNode, 1)
	queue[0] = root
	result := 0
	for len(queue) > 0 {
		size := len(queue)
		result = queue[0].Val
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
