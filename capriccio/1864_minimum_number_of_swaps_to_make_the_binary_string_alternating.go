package main

func minSwaps(s string) int {
	// There are only two cases.
	// First: 010101...
	// Second: 101010...
	FirstSwapToZero, FirstSwapToOne := 0, 0
	SecondSwapToZero, SencondSwapToOne := 0, 0
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] == '0' {
				SencondSwapToOne++
			} else if s[i] == '1' {
				FirstSwapToZero++
			}
		} else {
			if s[i] == '0' {
				FirstSwapToOne++
			} else if s[i] == '1' {
				SecondSwapToZero++
			}
		}
	}

	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	if FirstSwapToOne == FirstSwapToZero && SencondSwapToOne == SecondSwapToZero {
		return min(FirstSwapToOne, SencondSwapToOne)
	} else if FirstSwapToOne == FirstSwapToZero {
		return FirstSwapToOne
	} else if SencondSwapToOne == SecondSwapToZero {
		return SencondSwapToOne
	} else {
		return -1
	}
}
