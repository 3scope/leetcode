package capriccio

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	fast, slow := head.Next.Next, head.Next
	for fast != slow {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
