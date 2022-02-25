package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head
	fast, slow := newHead, newHead
	for i := 0; i < n; i++ {
		fast = fast.Next
		if fast == nil {
			return head
		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}
