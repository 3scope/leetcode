package main

import "container/heap"

type KthLargest struct {
	Data   []int
	Length int
	K      int
}

func (k KthLargest) Len() int {
	return k.Length
}

func (k KthLargest) Less(i, j int) bool {
	// 构建小根堆，“i”作为下标比“j”大，当“Data[i]”比“Data[j]”小的时候，交换两者位置。
	return k.Data[i] < k.Data[j]
}

func (k KthLargest) Swap(i, j int) {
	k.Data[i], k.Data[j] = k.Data[j], k.Data[i]
}

func (k *KthLargest) Push(x interface{}) {
	k.Data = append(k.Data, x.(int))
	k.Length++
}

func (k *KthLargest) Pop() interface{} {
	result := k.Data[k.Length-1]
	k.Data = k.Data[:k.Length-1]
	k.Length--
	return result
}

func ConstructorKth(k int, nums []int) KthLargest {
	h := KthLargest{
		Data: make([]int, 0, k+1),
		K:    k,
	}
	for i := 0; i < len(nums); i++ {
		h.Add(nums[i])
	}
	return h
}

func (k *KthLargest) Add(val int) int {
	heap.Push(k, val)
	if k.Length > k.K {
		heap.Pop(k)
	}
	result := heap.Pop(k).(int)
	heap.Push(k, result)

	return result
}
