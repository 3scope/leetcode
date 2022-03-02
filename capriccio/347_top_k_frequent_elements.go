package main

import (
	"container/heap"
	"sort"
)

// Priority queue, or called heap.
// Use an array to represent a complete binary tree.
// Each element of two-dimentional array contains value(at index of 0) and frequence(at index 0f 2).
type intHeap [][2]int

func (h intHeap) Len() int { return len(h) }

// The variable i is larger than j.
func (h intHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequentUseingHeap(nums []int, k int) []int {
	frequence := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		frequence[nums[i]]++
	}
	h := &intHeap{}
	heap.Init(h)
	for value, fre := range frequence {
		heap.Push(h, [2]int{value, fre})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = h.Pop().([2]int)[0]
	}
	return result
}

// Use sort method.
func topKFrequent(nums []int, k int) []int {
	frequence := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		frequence[nums[i]]++
	}
	result := make([]int, 0, len(frequence))
	for k, _ := range frequence {
		result = append(result, k)
	}
	sort.Slice(result, func(i, j int) bool {
		return frequence[result[i]] > frequence[result[j]]
	})

	return result[:k]
}

func main() {
	nums := []int{3, 0, 1, 0}
	topKFrequentUseingHeap(nums, 1)
}
