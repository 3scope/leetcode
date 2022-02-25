package main

type MyStack struct {
	queue *myQueue
}

type myQueue struct {
	Data   []int
	Length int
	Head   int
}

func (this *myQueue) Push(x int) {
	this.Length++
	this.Data = append(this.Data, x)
}

func (this *myQueue) Pop() int {
	this.Head++
	this.Length--
	return this.Data[this.Head-1]
}

func (this *myQueue) Empty() bool {
	if this.Length <= 0 {
		return true
	} else {
		return false
	}
}

func Constructor() MyStack {
	return MyStack{
		&myQueue{
			Data:   make([]int, 0),
			Length: 0,
			Head:   0,
		},
	}
}

func (this *MyStack) Push(x int) {
	this.queue.Push(x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < this.queue.Length-1; i++ {
		temp := this.queue.Pop()
		this.queue.Push(temp)
	}
	return this.queue.Pop()
}

func (this *MyStack) Top() int {
	result := this.Pop()
	this.Push(result)
	return result
}

func (this *MyStack) Empty() bool {
	return this.queue.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
