package main

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	curGas := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i] - cost[i]
		curGas += gas[i] - cost[i]
		if curGas < 0 {
			curGas = 0
			start = i + 1
		}
	}

	if totalGas < 0 {
		return -1
	}
	return start
}
