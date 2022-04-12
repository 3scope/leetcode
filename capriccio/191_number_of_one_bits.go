package main

func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i <= 31; i++ {
		bit := num >> i & 1
		if bit == 1 {
			count++
		}
	}
	return count
}
