package main

func robThree(root *TreeNode) int {
	result := robTree(root)
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	return max(result[0], result[1])
}

// The first element of slice is standing not steal this node, and the second is standing steal.
func robTree(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := robTree(root.Left)
	right := robTree(root.Right)

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	steal := root.Val + left[0] + right[0]
	notSteal := max(left[0], left[1]) + max(right[0], right[1])

	return []int{notSteal, steal}
}
