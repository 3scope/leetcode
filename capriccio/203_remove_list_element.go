package main

func removeElements(head *ListNode, val int) *ListNode {
	// 新建一个头节点，可以将所有节点统一处理。
	newHead := new(ListNode)
	newHead.Next = head
	pre := newHead
	for pre.Next != nil {
		// 相同就删除节点。
		if pre.Next.Val == val {
			cur := pre.Next
			pre.Next = cur.Next
		} else {
			// 不相同就向后遍历。
			pre = pre.Next
		}
	}

	return newHead.Next
}
