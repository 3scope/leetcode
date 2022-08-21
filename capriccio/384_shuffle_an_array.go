package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	Data []int
}

func ConstructorShuffle(nums []int) Solution {
	return Solution{
		Data: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.Data
}

func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	temp := make([]int, len(this.Data))
	copy(temp, this.Data)
	for i := len(temp) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		temp[i], temp[j] = temp[j], temp[i]
	}

	return temp
}
