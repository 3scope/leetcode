package main

import "math"

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	newRoot := new(TreeNode)
	newRoot.Val = int(math.Pow10(5)) + 1
	newRoot.Left = root
	pre, cur := newRoot, root
	for cur != nil {
		if cur.Val > key {
			pre = cur
			cur = cur.Left
		} else if cur.Val < key {
			pre = cur
			cur = cur.Right
		} else {
			break
		}
	}
	if cur == nil {
		return root
	}
	if cur.Left == nil && cur.Right == nil {
		if pre.Val > key {
			pre.Left = nil
		} else {
			pre.Right = nil
		}
	} else if cur.Left != nil && cur.Right == nil {
		if pre.Val > key {
			pre.Left = cur.Left
		} else {
			pre.Right = cur.Left
		}
	} else if cur.Left == nil && cur.Right != nil {
		if pre.Val > key {
			pre.Left = cur.Right
		} else {
			pre.Right = cur.Right
		}
	} else {
		rightChildleft := cur.Right
		for rightChildleft.Left != nil {
			rightChildleft = rightChildleft.Left
		}
		rightChildleft.Left = cur.Left
		if pre.Val > key {
			pre.Left = cur.Right
		} else {
			pre.Right = cur.Right
		}
	}

	return newRoot.Left
}

func deleteNodeRecursion(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left != nil && root.Right != nil {
			rightChildLeft := root.Right
			for rightChildLeft.Left != nil {
				rightChildLeft = rightChildLeft.Left
			}
			rightChildLeft.Left = root.Left
			return root.Right
		} else if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left != nil {
			return root.Left
		} else {
			return root.Right
		}
	}

	if root.Val > key {
		root.Left = deleteNodeRecursion(root.Left, key)
	} else {
		root.Right = deleteNodeRecursion(root.Right, key)
	}

	return root
}
