package main

type MyQueue struct {
	InputStack  *myStack
	OutputStack *myStack
}

type myStack struct {
	Data   []int
	Length int
}

func (this *myStack) Push(x int) {
	if cap(this.Data) > this.Length {
		this.Data[this.Length] = x
		this.Length++
	}
}

func (this *myStack) Pop() int {
	this.Length--
	return this.Data[this.Length]
}

func (this *myStack) Peek() int {
	return this.Data[this.Length-1]
}

func (this *myStack) Empty() bool {
	if this.Length <= 0 {
		return true
	} else {
		return false
	}
}

// func Constructor() MyQueue {
// 	return MyQueue{
// 		InputStack: &myStack{
// 			Data:   make([]int, 8),
// 			Length: 0,
// 		},
// 		OutputStack: &myStack{
// 			Data:   make([]int, 8),
// 			Length: 0,
// 		},
// 	}
// }

func (this *MyQueue) Push(x int) {
	this.InputStack.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.OutputStack.Empty() {
		for !this.InputStack.Empty() {
			temp := this.InputStack.Pop()
			this.OutputStack.Push(temp)
		}
	}
	return this.OutputStack.Pop()
}

func (this *MyQueue) Peek() int {
	result := this.Pop()
	this.OutputStack.Push(result)
	return result
}

func (this *MyQueue) Empty() bool {
	return this.InputStack.Empty() && this.OutputStack.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
