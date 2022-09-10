package main

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	subResult := make([]int, 0)
	combineBacktracking(n, k, 1, &result, &subResult)
	return result
}

// 变量“n, k, index”构成了回溯过程中，每一层的取值范围。
func combineBacktracking(n, k, index int, result *[][]int, subResult *[]int) {
	// 收集结果。
	if len(*subResult) == k {
		t := make([]int, k)
		copy(t, *subResult)
		*result = append(*result, t)
		return
	}

	// 剪枝操作。
	if (n-index+1)+len(*subResult) < k {
		return
	}

	for i := index; i <= n; i++ {
		*subResult = append(*subResult, i)
		combineBacktracking(n, k, i+1, result, subResult)
		// 进行回溯。
		*subResult = (*subResult)[:len(*subResult)-1]
	}
}
