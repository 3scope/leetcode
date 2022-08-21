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
	// 每当到叶子节点的时候，停止递归，记录路径。
	if root.Left == nil && root.Right == nil {
		// 把自身加入路径中。
		*path = append((*path), root)
		// 转换为字符串。
		pathString := strconv.Itoa((*path)[0].Val)
		for i := 1; i < len(*path); i++ {
			pathString += "->"
			pathString += strconv.Itoa((*path)[i].Val)
		}
		// 存储结果。
		*result = append((*result), pathString)
	}
	// 回溯法。
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
