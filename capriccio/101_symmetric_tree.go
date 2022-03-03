package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 2)
	queue[0] = root.Left
	queue[1] = root.Right
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i = i + 2 {
			rootLeft := queue[0]
			rootRight := queue[1]
			queue = queue[2:]
			if rootLeft == nil && rootRight == nil {
				continue
			} else if rootLeft == nil || rootRight == nil || rootLeft.Val != rootRight.Val {
				return false
			}

			queue = append(queue, rootLeft.Left, rootRight.Right, rootLeft.Right, rootRight.Left)
		}
	}
	return true
}

func main() {
	nodeIntSlice := []int{1, 2, 2, 3, 4, 4, 3}
	node := make([]*TreeNode, len(nodeIntSlice))
	for i := 0; i < len(nodeIntSlice); i++ {
		cur := new(TreeNode)
		cur.Val = nodeIntSlice[i]
		node[i] = cur
	}
	// Use node slice to build tree.
	for i := 0; i < len(node); i++ {
		if 2*i+1 >= len(node) {
			break
		}
		node[i].Left = node[2*i+1]
		if 2*i+2 < len(node) {
			node[i].Right = node[2*i+2]
		}
	}
	isSymmetric(node[0])
}
