package main

func lastStoneWeightII(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}

	target := sum / 2
	dp := make([]int, target+1)
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	// The weight of other is greater than target.
	other := sum - dp[target]
	return other - dp[target]
}
