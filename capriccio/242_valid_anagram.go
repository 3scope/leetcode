package capriccio

func isAnagram(s string, t string) bool {
	flag := true
	dictionary := make(map[rune]int)
	for _, char := range s {
		dictionary[char]++
	}
	for _, char := range t {
		dictionary[char]--
	}
	for key, _ := range dictionary {
		if dictionary[key] != 0 {
			flag = false
		}
	}
	return flag
}
