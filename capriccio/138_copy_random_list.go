package main

type newNode struct {
	Val    int
	Next   *newNode
	Random *newNode
}

func copyRandomList(head *newNode) *newNode {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		temp := new(newNode)
		temp.Val = node.Val
		temp.Next = node.Next
		node.Next = temp
	}
	for node := head; node != nil; node = node.Next.Next {
		temp := node.Next
		if node.Random != nil {
			temp.Random = node.Random.Next
		}
	}

	// Remove the redundant next pointer.
	newHead := head.Next
	for node, old := head.Next, head; ; {
		if node.Next != nil {
			third := node.Next
			node.Next = third.Next
			node = node.Next

			old.Next = third
			old = old.Next
		} else {
			old.Next = nil
			break
		}
	}

	return newHead
}

func copyRandomListTwo(head *newNode) *newNode {
	if head == nil {
		return nil
	}
	newHead := new(newNode)
	newHead.Val = head.Val
	// The map is to store the relation of old list and new list.
	nodeMap := make(map[*newNode]*newNode)
	nodeMap[head] = newHead

	for node, temp := head.Next, newHead; node != nil; node = node.Next {
		cur := new(newNode)
		cur.Val = node.Val
		temp.Next = cur
		temp = cur
		nodeMap[node] = temp
	}

	for node, cur := head, newHead; node != nil; node, cur = node.Next, cur.Next {
		if node.Random != nil {
			cur.Random = nodeMap[node.Random]
		}
	}

	return newHead
}
