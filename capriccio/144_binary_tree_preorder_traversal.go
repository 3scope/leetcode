package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 1)
	// 定义函数。
	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		// 终止条件。
		if root == nil {
			return
		}
		// 先对当前节点进行处理。
		result = append(result, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	// 执行函数
	preorder(root)

	return result
}

func preorderTraversalIteration(root *TreeNode) []int {
	// 特殊情况。
	if root == nil {
		return nil
	}
	result := make([]int, 0, 1)
	// 使用栈来模拟树的递归。
	stack := make([]*TreeNode, 1)
	stack[0] = root
	// 空节点不会入栈，所以只要栈不空，就不会结束循环。
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 对当前节点处理，满足先序遍历。
		result = append(result, cur.Val)
		// 先入栈右节点，后入栈左节点，保证栈的顺序。
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return result
}
