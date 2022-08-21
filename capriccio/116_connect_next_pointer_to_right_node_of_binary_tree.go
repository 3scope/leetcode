package main

type NextNode struct {
	Val   int
	Left  *NextNode
	Right *NextNode
	Next  *NextNode
}

func connect(root *NextNode) *NextNode {
	if root == nil {
		return nil
	}
	// 使用层序遍历，遍历中每层的前一个节点指向后一个，最后一个的Next域置为空。
	queue := make([]*NextNode, 1)
	queue[0] = root
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size-1; i++ {
			cur := queue[0]
			queue = queue[1:]
			cur.Next = queue[0]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		last := queue[0]
		queue = queue[1:]
		if last.Left != nil {
			queue = append(queue, last.Left)
		}
		if last.Right != nil {
			queue = append(queue, last.Right)
		}
	}

	return root
}
