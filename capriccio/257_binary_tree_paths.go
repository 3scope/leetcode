package main

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	path := make([]*TreeNode, 0)
	result := make([]string, 0)
	traversalRecursion(root, &path, &result)

	return result
}

func traversalRecursion(root *TreeNode, path *[]*TreeNode, result *[]string) {
	if root.Left == nil && root.Right == nil {
		// Add node itself.
		*path = append((*path), root)

		pathString := strconv.Itoa((*path)[0].Val)
		for i := 1; i < len(*path); i++ {
			pathString += "->"
			pathString += strconv.Itoa((*path)[i].Val)
		}
		*result = append((*result), pathString)
	}

	if root.Left != nil {
		*path = append((*path), root)
		traversalRecursion(root.Left, path, result)
		*path = (*path)[:len(*path)-1]
	}
	if root.Right != nil {
		*path = append((*path), root)
		traversalRecursion(root.Right, path, result)
		*path = (*path)[:len(*path)-1]
	}
}
