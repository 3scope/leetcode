package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	temp := make([]string, 0)
	restoreIpAddressesBackTracking(s, 0, &result, &temp)

	return result
}

func restoreIpAddressesBackTracking(s string, startIndex int, result *[]string, temp *[]string) {
	if startIndex >= len(s) && len(*temp) == 4 {
		*result = append(*result, strings.Join(*temp, "."))
		return
	}

	// The max length of each segment is 3, and i should less than the origin string length.
	for i := startIndex; i < startIndex+3 && i < len(s); i++ {
		if islegalIPAddress(s[startIndex : i+1]) {
			*temp = append(*temp, s[startIndex:i+1])
			restoreIpAddressesBackTracking(s, i+1, result, temp)
			*temp = (*temp)[:len(*temp)-1]
		}
	}
}

func islegalIPAddress(s string) bool {
	num, _ := strconv.Atoi(s)
	if num <= 255 && (s[0] != '0' || len(s) == 1) {
		return true
	}
	return false
}
