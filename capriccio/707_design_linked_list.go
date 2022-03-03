package main

type MyLinkedList struct {
	Head   *ListNode
	Tail   *ListNode
	Length int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// func Constructor() MyLinkedList {
// 	return MyLinkedList{
// 		Head:   new(ListNode),
// 		Tail:   nil,
// 		Length: 0,
// 	}
// }

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Length || index < 0 {
		return -1
	}
	node := this.Head
	for ; index >= 0; index-- {
		node = node.Next
	}
	return node.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := new(ListNode)
	node.Val = val
	node.Next = this.Head.Next
	this.Head.Next = node
	if this.Length == 0 {
		this.Tail = node
	}
	this.Length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := new(ListNode)
	node.Val = val
	if this.Length == 0 {
		this.AddAtHead(val)
	} else {
		this.Tail.Next = node
		this.Tail = node
		this.Length++
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.Length {
		this.AddAtTail(val)
	} else if index < this.Length && index >= 0 {
		node := new(ListNode)
		node.Val = val
		pre := this.Head
		for ; index > 0; index-- {
			pre = pre.Next
		}
		node.Next = pre.Next
		pre.Next = node
		this.Length++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < this.Length && index >= 0 {
		pre := this.Head
		for ; index > 0; index-- {
			pre = pre.Next
		}

		if index == (this.Length - 1) {
			this.Tail = pre
			pre.Next = nil
		} else {
			pre.Next = pre.Next.Next
		}
		this.Length--
	}
}
