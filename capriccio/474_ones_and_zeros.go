package main

func findMaxForm(strs []string, m int, n int) int {
	// The value of m stand for count of "0".
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		zero := 0
		one := 0
		for _, v := range str {
			if v == '0' {
				zero++
			} else if v == '1' {
				one++
			}
		}
		for j := m; j >= zero; j-- {
			for k := n; k >= one; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zero][k-one]+1)
			}
		}
	}
	return dp[m][n]
}
