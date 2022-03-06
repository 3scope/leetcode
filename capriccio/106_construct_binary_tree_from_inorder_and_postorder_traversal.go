package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	root := new(TreeNode)
	root.Val = val
	var leftInorder, rightInorder []int
	for i := 0; i < len(inorder); i++ {
		if val == inorder[i] {
			leftInorder = inorder[:i]
			rightInorder = inorder[i+1:]
		}
	}
	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder) : len(postorder)-1]
	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)

	return root
}
