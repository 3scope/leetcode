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

type MonotonicQueue struct {
	// The front of the queue is at index of 0.
	Data   []int
	Length int
}

func (mq *MonotonicQueue) Front() int {
	if mq.Length > 0 {
		return mq.Data[0]
	} else {
		return -1
	}
}

func (mq *MonotonicQueue) Pop(value int) {
	if mq.Length > 0 && mq.Data[0] == value {
		mq.PopFront()
	}
}

func (mq *MonotonicQueue) PopBack() {
	if mq.Length > 0 {
		mq.Data = mq.Data[:mq.Length-1]
		mq.Length--
	}
}

func (mq *MonotonicQueue) PopFront() {
	if mq.Length > 0 {
		mq.Data = mq.Data[1:mq.Length]
		mq.Length--
	}
}

func (mq *MonotonicQueue) Push(value int) {
	// Monotonic non-decreasing queue.
	for mq.Length != 0 && mq.Data[mq.Length-1] < value {
		mq.PopBack()
	}
	mq.Data = append(mq.Data, value)
	mq.Length++
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := &MonotonicQueue{
		Data:   make([]int, 0, 8),
		Length: 0,
	}
	// To initialize the monotonic queue.
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	// The first round.
	result := []int{queue.Front()}
	for i := k; i < len(nums); i++ {
		queue.Pop(nums[i-k])
		queue.Push((nums[i]))
		result = append(result, queue.Front())
	}

	return result
}
