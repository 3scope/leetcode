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
	*path = append((*path), root)
	if root.Left == nil && root.Right == nil {
		pathString := strconv.Itoa((*path)[0].Val)
		for i := 1; i < len(*path); i++ {
			pathString += "->"
			pathString += strconv.Itoa((*path)[i].Val)
		}
		*result = append((*result), pathString)
	}

	if root.Left != nil {
		traversalRecursion(root.Left, path, result)
		*path = (*path)[:len(*path)-1]
	}
	if root.Right != nil {
		traversalRecursion(root.Right, path, result)
		*path = (*path)[:len(*path)-1]
	}
}
