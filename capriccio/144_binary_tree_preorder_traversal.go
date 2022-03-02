package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	var preorder func(*TreeNode, *[]int)
	preorder = func(root *TreeNode, result *[]int) {
		if root == nil {
			return
		}
		*result = append(*result, root.Val)
		preorder(root.Left, result)
		preorder(root.Right, result)
	}
	preorder(root, &result)

	return result
}

func preorderTraversalUsingIteration(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 1)
	stack[0] = root
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return result
}
