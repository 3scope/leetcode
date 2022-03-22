package main

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	temp := make([]int, 0)
	combinationSum3Backtracking(k, n, 0, 1, &result, &temp)

	return result
}

func combinationSum3Backtracking(k, n, sum, startIndex int, result *[][]int, temp *[]int) {
	if len(*temp) == k {
		if sum == n {
			new := make([]int, k)
			for i := 0; i < k; i++ {
				new[i] = (*temp)[i]
			}
			*result = append(*result, new)
		}
		return
	}

	if sum >= n {
		return
	}

	for i := startIndex; i <= 9; i++ {
		sum += i
		*temp = append(*temp, i)
		combinationSum3Backtracking(k, n, sum, i+1, result, temp)
		sum -= i
		*temp = (*temp)[:len(*temp)-1]
	}
}
