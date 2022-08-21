package main

func reverseList(head *ListNode) *ListNode {
	// 特殊情况。
	if head == nil {
		return head
	}
	// 定义两个指针，一个指向前驱节点，一个指向当前节点。
	var pre, cur *ListNode
	pre = nil
	cur = head
	for cur != nil {
		// 存储“cur”的下一个节点。
		temp := cur.Next
		// 翻转的过程。
		cur.Next = pre
		// 更新指针。
		pre = cur
		cur = temp
	}
	return pre
}
