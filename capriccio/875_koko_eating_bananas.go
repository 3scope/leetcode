package main

// 可以使用二分查找的方式，最主要的点是计算某个速度下吃一堆香蕉需要的时间。
func minEatingSpeed(piles []int, h int) int {
	// 最慢的速度。
	left := 1
	// 最快的速度取决于某一堆香蕉最多有多少个。
	right := 0
	for i := 0; i < len(piles); i++ {
		if right < piles[i] {
			right = piles[i]
		}
	}
	// 二分查找。
	middle := left + (right-left)/2
	for left <= right {
		if getTimeOfPiles(piles, middle) < h {
			// 吃的太快，需要减小“middle”的值。
			right = middle - 1
		} else if getTimeOfPiles(piles, middle) > h {
			// 吃的太慢，需要增加“middle”的值。
			left = middle + 1
		} else {
			// 如果满足要求的话，那么尝试减少吃的速度。
			right = middle - 1
		}
		middle = left + (right-left)/2
	}
	return middle
}

// 计算吃完所有香蕉的时间。
func getTimeOfPiles(piles []int, k int) int {
	time := 0
	for i := 0; i < len(piles); i++ {

		time += piles[i] / k
		if piles[i]%k != 0 {
			time++
		}
	}
	return time
}
