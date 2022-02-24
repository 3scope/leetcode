package capriccio

func reverseStr(s string, k int) string {
	result := []rune(s)
	for i := 0; i < len(s); i += 2 * k {
		var start, end int
		if len(s)-i >= k {
			start = i
			end = i + k - 1
		} else {
			start = i
			end = len(s) - 1
		}
		for ; start < end; start, end = start+1, end-1 {
			result[start], result[end] = result[end], result[start]
		}
	}
	return string(result)
}
