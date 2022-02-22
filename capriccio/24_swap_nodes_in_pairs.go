package capriccio

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := new(ListNode)
	newHead.Next = head
	current := newHead
	var currentNext, current3Next *ListNode
	for current.Next != nil && current.Next.Next != nil {
		currentNext = current.Next
		current3Next = current.Next.Next.Next
		// The pointer filed of current points to next next.
		current.Next = current.Next.Next
		current.Next.Next = currentNext
		currentNext.Next = current3Next
		current = currentNext
	}
	return newHead.Next
}
