package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

// OverFlow
func addTwoNumbersOverFlow(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	count := 0
	num1 := 0
	num2 := 0

	for l1 != nil {
		num1 += l1.Val * int(math.Pow10(count))
		count++
		l1 = l1.Next
	}

	count = 0
	for l2 != nil {
		num2 += l2.Val * int(math.Pow10(count))
		count++
		l2 = l2.Next
	}
	// Right sum.
	sum := num1 + num2

	if sum < 10 {
		result.Val = sum
		result.Next = nil
		return result
	} else {
		result.Val = sum % 10
		result.Next = nil
		sum = sum / 10
	}

	var ptr *ListNode = result
	for sum != 0 {
		temp := new(ListNode)
		temp.Val = sum % 10
		ptr.Next = temp
		ptr = temp
		sum = sum / 10
	}

	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// The head node.
	result := new(ListNode)
	// The current node pointer.
	current := result

	carry := 0

	for l1 != nil || l2 != nil {
		var1, var2 := 0, 0
		if l1 != nil {
			var1 = l1.Val
		}

		if l2 != nil {
			var2 = l2.Val
		}

		value := var1 + var2 + carry

		temp := new(ListNode)
		temp.Val = value % 10
		carry = value / 10
		current.Next = temp
		current = temp

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry == 1 {
		temp := new(ListNode)
		temp.Val = 1
		current.Next = temp
		current = temp
	}

	return result.Next
}
