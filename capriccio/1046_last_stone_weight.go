package main

import "container/heap"

type MaxIntHeap struct {
	Data   []int
	Length int
}

func (h MaxIntHeap) Len() int {
	return h.Length
}

func (h MaxIntHeap) Less(i, j int) bool {
	return h.Data[i] > h.Data[j]
}

func (h MaxIntHeap) Swap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

func (h *MaxIntHeap) Push(x interface{}) {
	h.Data = append(h.Data, x.(int))
	h.Length++
}

func (h *MaxIntHeap) Pop() interface{} {
	result := h.Data[h.Length-1]
	h.Data = h.Data[:h.Length-1]
	h.Length--
	return result
}

func lastStoneWeight(stones []int) int {
	h := &MaxIntHeap{
		Data:   stones,
		Length: len(stones),
	}
	heap.Init(h)
	for h.Len() > 1 {
		left := heap.Pop(h).(int)
		right := heap.Pop(h).(int)
		newStone := func() int {
			if left > right {
				return left - right
			} else {
				return right - left
			}
		}()
		if newStone > 0 {
			heap.Push(h, newStone)
		}
	}
	if h.Len() == 1 {
		return h.Data[0]
	}
	return 0
}
