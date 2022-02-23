package capriccio

func isHappy(n int) bool {
	if getSum(n) == 1 || n == 1 {
		return true
	}
	duplication := make(map[int]bool)
	duplication[n] = true

	for getSum(n) != 1 {
		n = getSum(n)
		if duplication[n] {
			return false
		} else {
			duplication[n] = true
		}
	}
	return true
}

func getSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
