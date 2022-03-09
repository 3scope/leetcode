package main

func numTrees(n int) int {
	dp := make([]int, n+1)
	// When the tree is nil, it's still a binary search tree.
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// The "j-1" is the number of left node, and the "i-j" is the number of right node.
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
