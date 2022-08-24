package main

// 设计带有头尾指针的链表，并且提供链表长度，可以在O(1)的时间复杂度得到链表长度。
type MyLinkedList struct {
	Head   *ListNode
	Length int
}

// 单链表节点，LeetCode上报错可能是重复定义。
type ListNode struct {
	Val  int
	Next *ListNode
}

// 通过下标得到节点值，下标从0开始。
func (this *MyLinkedList) Get(index int) int {
	// 下标索引无效就返回-1。
	if index >= this.Length || index < 0 {
		return -1
	}
	// “node”从头开始遍历。
	node := this.Head
	for i := index; i >= 0; i-- {
		node = node.Next
	}
	return node.Val
}

// 头插法添加节点。
func (this *MyLinkedList) AddAtHead(val int) {
	// 特殊情况，如果链表为空，则使用尾插法。
	if this.Length == 0 {
		this.AddAtTail(val)
		return
	}
	// 生成新的节点。
	node := new(ListNode)
	node.Val = val
	// 头插法。
	node.Next = this.Head.Next
	this.Head.Next = node
	this.Length++
}

// 尾插法添加节点，尾插法会更新 Tail 指针。
func (this *MyLinkedList) AddAtTail(val int) {
	// 生成新的节点。
	node := new(ListNode)
	node.Val = val
	// 尾插法更新 Tail 指针。
	pre := this.Head
	// 寻找索引值为“Length-1”的节点，即最后一个节点。
	for i := this.Length - 1; i >= 0; i-- {
		pre = pre.Next
	}
	pre.Next = node
	this.Length++
}

// 在某个节点前插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.Length {
		// 如果“index”等于链表长度，则相当于尾插法。
		this.AddAtTail(val)
	} else if index < this.Length && index >= 0 {
		// 生成节点。
		node := new(ListNode)
		node.Val = val
		// 查找前驱节点。
		pre := this.Head
		for i := index; i > 0; i-- {
			pre = pre.Next
		}
		node.Next = pre.Next
		pre.Next = node
		this.Length++
	}
	// 如果“index”大于链表长度，则不插入节点。
}

// 删除某个指定下标的节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < this.Length && index >= 0 {
		// 找到删除节点的前驱节点。
		pre := this.Head
		for i := index; i > 0; i-- {
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
		this.Length--
	}
	// 如果“index”无效，则不删除节点。
}
