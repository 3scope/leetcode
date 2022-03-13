package main

func maxProfitWithFee(prices []int, fee int) int {
	if len(prices) == 1 {
		return 0
	}
	result := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i] >= minPrice && prices[i] <= minPrice+fee {
			continue
		} else {
			result += prices[i] - fee - minPrice
			// The profit of the next round is the difference between the current price and the next one.
			minPrice = prices[i] - fee
		}
	}

	return result
}
