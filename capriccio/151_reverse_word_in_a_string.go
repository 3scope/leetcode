package capriccio

import "strings"

func reverseWords(s string) string {
	temp := strings.Split(s, " ")
	noNullTemp := make([]string, 0)
	for _, value := range temp {
		if value != "" {
			noNullTemp = append(noNullTemp, value)
		}
	}
	for i, j := 0, len(noNullTemp)-1; i < j; i, j = i+1, j-1 {
		noNullTemp[i], noNullTemp[j] = noNullTemp[j], noNullTemp[i]
	}
	result := strings.Join(noNullTemp, " ")

	return result
}

// Use byte slice to reverse.
func reverseWordsInSentance(s string) string {
	if s == "" {
		return ""
	}
	temp := []byte(s)
	for len(temp) != 0 && temp[0] == ' ' {
		temp = temp[1:]
	}
	for len(temp) != 0 && temp[len(temp)-1] == ' ' {
		temp = temp[:len(temp)-1]
	}
	if len(temp) == 0 {
		return ""
	}
	for i := 1; i < len(temp); i++ {
		if temp[i] == temp[i-1] && temp[i] == ' ' {
			temp = append(temp[:i-1], temp[i:]...)
			i--
		}
	}
	// Totally reverse.
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	// Reverse words.
	for i := 0; i < len(temp); {
		j := i
		for ; j < len(temp); j++ {
			if temp[j] == ' ' {
				break
			}
		}
		start := j
		j--
		for ; i < j; i, j = i+1, j-1 {
			temp[i], temp[j] = temp[j], temp[i]
		}
		i = start + 1
	}
	return string(temp)
}
