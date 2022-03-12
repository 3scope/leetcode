package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// To add a head pointer.
	newHead := new(ListNode)
	newHead.Next = head

	temp := newHead
	length := 0
	// To get the length of the list.
	for temp.Next != nil {
		length++
		temp = temp.Next
	}

	k = k % length
	// If k equals to 0, do nothing.
	if k == 0 {
		return head
	}

	slow, fast := newHead, newHead
	for i := 0; i < length; i++ {
		if i >= k {
			slow = slow.Next
		}
		fast = fast.Next
	}
	newHead.Next = slow.Next
	slow.Next = nil
	fast.Next = head

	return newHead.Next
}
