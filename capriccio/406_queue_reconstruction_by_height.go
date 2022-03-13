package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	if len(people) == 1 {
		return people
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return false
	})

	for i := 0; i < len(people); i++ {
		if i != people[i][1] {
			index := people[i][1]
			temp := people[i]
			copy(people[index+1:], people[index:i])
			people[index] = temp
		}
	}

	return people
}
