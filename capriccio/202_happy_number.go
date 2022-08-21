package main

func isHappy(n int) bool {
	// 如果本身是1，或者一次计算就得到1，那么它就是快乐数。
	if getSum(n) == 1 || n == 1 {
		return true
	}
	duplication := make(map[int]bool)
	duplication[n] = true

	for getSum(n) != 1 {
		n = getSum(n)
		// 如果出现重复，那么再次计算就会循环，这时直接返回“false”。
		if duplication[n] {
			return false
		} else {
			duplication[n] = true
		}
	}
	return true
}

// 得到每位数的平方和。
func getSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
