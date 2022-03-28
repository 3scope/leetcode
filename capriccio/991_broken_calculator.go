package main

func brokenCalc(startValue int, target int) int {
	if target <= startValue {
		return startValue - target
	}

	// Reverse thinking, target approaches startValue.
	// If target greater than startValue, the only thing it can do is to divede 2, and if it is odd, it should plus one.
	result := 0
	for startValue < target {
		if target%2 == 0 {
			target = target / 2
		} else {
			target = target + 1
		}
		result++
	}

	return result
}
