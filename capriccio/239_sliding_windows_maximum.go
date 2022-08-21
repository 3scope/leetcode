package main

import "sort"

func bruteForceMaxSlidingWindow(nums []int, k int) []int {
	result := make([]int, len(nums)-k+1)
	for i := 0; i <= len(nums)-k; i++ {
		temp := make([]int, k)
		copy(temp, nums[i:i+k])
		sort.Ints(temp)
		result[i] = temp[k-1]
	}

	return result
}

// 采用单调队列的方式，队列中队头元素到队尾元素为单调非增的。
type MonotonicQueue struct {
	Data   []int
	Length int
}

// 得到队头元素。
func (mq *MonotonicQueue) Front() int {
	if mq.Length > 0 {
		return mq.Data[0]
	} else {
		return -1
	}
}

// 从队头推出元素。
func (mq *MonotonicQueue) Pop(value int) {
	if mq.Length > 0 && mq.Data[0] == value {
		mq.Data = mq.Data[1:mq.Length]
		mq.Length--
	}
}

func (mq *MonotonicQueue) Push(value int) {
	// 当新增一个元素的时候，需要将队尾中比它小的元素推出。
	for mq.Length != 0 && mq.Data[mq.Length-1] < value {
		mq.Data = mq.Data[:mq.Length-1]
		mq.Length--
	}
	mq.Data = append(mq.Data, value)
	mq.Length++
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := &MonotonicQueue{
		Data:   make([]int, 0, 8),
		Length: 0,
	}
	// 初始化第一个窗口。
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	// 存储第一个窗口的结果，之后每次移动，都存储新的结果。
	result := []int{queue.Front()}
	for i := k; i < len(nums); i++ {
		queue.Pop(nums[i-k])
		queue.Push((nums[i]))
		result = append(result, queue.Front())
	}
	return result
}
