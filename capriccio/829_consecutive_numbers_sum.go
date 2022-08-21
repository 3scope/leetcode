package main

// 由于求的是连续的数累加和等于“n”，并且从1开始，那么通过等差数列通项公式可以得到起始的数“a”和“n”的关系为：
//  (a + a + k - 1) * k / 2 = n => 2n / k = 2a + k - 1 = >  (2n / k) - k = 2a - 1 >= 1 => 2n / k > k
// 因为“2a - 1”和“k”都是整数，因此“k”一定是“2n”的因数。
func consecutiveNumbersSum(n int) int {
	result := 0
	for k := 1; k*k < 2*n; k++ {
		// “k”必须是“2n”的因数。
		if (2*n)%k != 0 {
			continue
		}
		// (a + a + k - 1) * k / 2 = n => (2n / k) - k + 1 = 2a => (2n / k - k + 1) % 2 == 0
		if (2*n/k-k+1)%2 == 0 {
			result++
		}
	}
	return result
}
