package main

// 求二叉树的深度最简单的方式是使用层序遍历，这样可以快速求出深度。
// 使用递归的话，有两种方式，第一种是求高度。
func maxDepthUsingHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 后序遍历，相当于求高度。
	left := maxDepthUsingHeight(root.Left)
	right := maxDepthUsingHeight(root.Right)
	// 遍历完子树后，可以选一个最高的子树，再加一。
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// 第二种方式是求深度，可以使用回溯法。
func maxDepthUsingDepth(root *TreeNode, depth int, result *int) {
	// 终止条件。
	if root == nil {
		// 存一个最大值。
		if depth > *result {
			*result = depth
		}
		return
	}

	// 先序遍历。
	depth += 1
	maxDepthUsingDepth(root.Left, depth, result)
	maxDepthUsingDepth(root.Right, depth, result)
	depth -= 1
}
