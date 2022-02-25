package main

// func strStr(haystack string, needle string) int {
//     return strings.Index(haystack, needle)
// }

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	//整段判断是否相同
	//判断的时候需要注意数组越界，所以小于haystack长度减去needle的长度
	//判断的时候也需要注意两个长度是一样的
	needleLength := len(needle)
	for i := 0; i <= len(haystack)-needleLength; i++ {
		if haystack[i:i+needleLength] == needle {
			return i
		}
	}
	return -1
}
