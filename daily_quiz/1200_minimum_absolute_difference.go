package main

import (
	"math"
	"sort"
)

// func main() {
// 	arr := []int{100, 1, 2}
// 	result := minimumAbsDifference(arr)
// 	fmt.Println(result)
// }

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	// Actually the result value is [][2]int.
	result := make([][]int, 0)
	minimum := int(math.Pow10(6)) * 2

	// To append the appropriate result.
	for i := 1; i < len(arr); i++ {
		if minimum, ok := min(arr[i]-arr[i-1], minimum); ok {
			// Clear.
			result = result[0:0]
			// New value.
			element := []int{arr[i-1], arr[i]}
			// New result
			result = append(result, element)
		} else if (arr[i] - arr[i-1]) == minimum {
			element := []int{arr[i-1], arr[i]}
			result = append(result, element)
		} else {
			continue
		}
	}

	return result
}

func min(x, y int) (int, bool) {
	if x < y {
		return x, true
	} else {
		return y, false
	}
}
