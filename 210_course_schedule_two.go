package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	hashMap := make(map[int][]int)
	// First to set the requires.
	for i := 0; i < len(prerequisites); i++ {
		hashMap[prerequisites[i][0]] = append(hashMap[prerequisites[i][0]], prerequisites[i][1])
	}

	// Use hash map to set order.
	result := make([]int, 0)
	selected := make(map[int]int)
	for course := 0; course < numCourses; course++ {
		// Jump to next turn if course is in result.
		if _, ok := selected[course]; ok {
			continue
		}
		// Else to insert.
		result = setPre(course, result, hashMap, selected)
	}

	return result
}

func setPre(course int, result []int, hashMap map[int][]int, selected map[int]int) []int {
	if _, ok := selected[course]; ok {
		return result
	}
	selected[course]++

	if hashMap[course] == nil {
		result = append(result, course)
	} else {
		for _, value := range hashMap[course] {
			result = setPre(value, result, hashMap, selected)
		}
		result = append(result, course)
	}

	return result
}
