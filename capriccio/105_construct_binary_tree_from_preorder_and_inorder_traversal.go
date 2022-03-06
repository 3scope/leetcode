package main

func buildTreePre(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	root := new(TreeNode)
	root.Val = val

	var leftInorder, rightInorder []int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == val {
			leftInorder = inorder[:i]
			rightInorder = inorder[i+1:]
		}
	}
	leftPreorder := preorder[1 : len(leftInorder)+1]
	rigthPreorder := preorder[len(leftInorder)+1:]

	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rigthPreorder, rightInorder)

	return root
}
