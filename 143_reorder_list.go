package main

func reorderList(head *ListNode) {
	// Get the mid
	fast, slow := head, head
	if fast == nil || fast.Next == nil {
		return
	}
	// For break the chain.
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// Reverse.
	mid := slow.Next
	slow.Next = nil

	slice := make([]*ListNode, 0)
	for mid != nil {
		slice = append(slice, mid)
		mid = mid.Next
	}
	for i := len(slice) - 1; i > 0; i-- {
		slice[i].Next = slice[i-1]
		slice[i-1].Next = nil
	}
	mid = slice[len(slice)-1]

	// Merge.
	currentH := head
	for mid != nil {
		temp := currentH.Next
		currentH.Next = mid
		currentH = temp
		temp = mid.Next
		mid.Next = currentH
		mid = temp
	}
}
