package main

func canConstruct(ransomNote string, magazine string) bool {
	store := make(map[rune]int)
	for _, value := range ransomNote {
		store[value]++
	}
	for _, value := range magazine {
		store[value]--
	}
	for _, value := range store {
		if value > 0 {
			return false
		}
	}
	return true
}
