package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func TreeConstructor() Codec {
	return Codec{}
}

// 使用层序遍历，将一个二叉树变成一个字符串。
func (c *Codec) serialize(root *TreeNode) string {
	// 特殊情况。
	if root == nil {
		return "#"
	}

	result := ""
	// 使用队列来模拟层序。
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 空指针特殊处理。
			if cur == nil {
				result += "#-"
			} else {
				result += strconv.Itoa(cur.Val) + "-"
				// 空指针可以入队，但是会特殊处理。
				queue = append(queue, cur.Left)
				queue = append(queue, cur.Right)
			}
		}
	}
	// 末尾会多出一个“-”符号，所以需要删掉。
	result = result[:len(result)-1]

	return result
}

// 同样使用层序遍历，将字符串恢复为二叉树。
func (c *Codec) deserialize(data string) *TreeNode {
	// 特殊情况。
	if data == "#" {
		return nil
	}

	nodeStr := strings.Split(data, "-")
	nodeArray := make([]*TreeNode, len(nodeStr))
	// 创建节点。
	for i := 0; i < len(nodeArray); i++ {
		// 如果是空节点。
		if nodeStr[i] == "#" {
			nodeArray[i] = nil
		} else {
			nodeArray[i] = new(TreeNode)
			nodeArray[i].Val, _ = strconv.Atoi(nodeStr[i])
		}
	}

	// 层序遍历。
	queue := make([]*TreeNode, 1)
	queue[0] = nodeArray[0]
	// 变量 Index 指向遍历节点的孩子节点。
	index := 1
	// 按照顺序遍历 NodeArray 数组就是层序遍历。
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if nodeArray[index] != nil {
			cur.Left = nodeArray[index]
			queue = append(queue, cur.Left)
		}
		if nodeArray[index+1] != nil {
			cur.Right = nodeArray[index+1]
			queue = append(queue, cur.Right)
		}
		index += 2
	}

	return nodeArray[0]
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
