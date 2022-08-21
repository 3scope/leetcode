package main

import "sort"

func maximumSwap(num int) int {
	numArray := make([]int, 0)
	for num != 0 {
		numArray = append(numArray, num%10)
		num = num / 10
	}

	// 存放理论最大值，即所有数字从大到小排序。
	maxNum := make([]int, len(numArray))
	copy(maxNum, numArray)
	sort.Ints(maxNum)

	// 查看原字符串和理论最大值的区别，从高位开始处理。
	for i := len(numArray) - 1; i >= 0; i-- {
		// 处理第一个不同的位置。
		if numArray[i] != maxNum[i] {
			// 找到之后最大的位置。
			index := i - 1
			for j := i - 2; j >= 0; j-- {
				if numArray[j] >= numArray[index] {
					index = j
				}
			}
			numArray[i], numArray[index] = numArray[index], numArray[i]
			break
		}
	}

	result := 0
	for i := len(numArray) - 1; i >= 0; i-- {
		result = result*10 + numArray[i]
	}

	return result
}
