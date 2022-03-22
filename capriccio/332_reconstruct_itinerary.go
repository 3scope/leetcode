package main

import "sort"

func findItinerary(tickets [][]string) []string {
	if len(tickets) == 1 {
		return tickets[0]
	}

	// The plan struct is "source -> destination : number of flights".
	plan := make(map[string]map[string]int)
	count := 1
	for i := 0; i < len(tickets); i++ {
		if m, ok := plan[tickets[i][0]]; !ok {
			m = make(map[string]int)
			plan[tickets[i][0]] = m
		}
		plan[tickets[i][0]][tickets[i][1]]++
		count++
	}

	start := "JFK"
	result := []string{"JFK"}
	dfsItinerary(&plan, count, start, &result)

	return result
}

func dfsItinerary(plan *map[string]map[string]int, count int, start string, result *[]string) bool {
	if len(*result) == count {
		return true
	}

	if len((*plan)[start]) > 0 {
		dest := (*plan)[start]
		// Get all the destinations name.
		keys := make([]string, 0)
		for k, v := range dest {
			if v > 0 {
				keys = append(keys, k)
			}
		}
		// Sorting keys.
		sort.Strings(keys)
		// Then do traversal.
		for i := 0; i < len(keys); i++ {
			dest[keys[i]] -= 1
			*result = append(*result, keys[i])
			flag := dfsItinerary(plan, count, keys[i], result)
			if flag {
				return true
			}
			dest[keys[i]] += 1
			*result = (*result)[:len(*result)-1]
		}
	}

	// If the destination has no itinerary to other station, backtracking.
	return false
}

func dfsItineraryIteration(plan map[string]map[string]int, count int, start string, result *[]string) {
	stack := []string{start}
	for len(stack) > 0 {
		size := len(stack)
		// Search for every level.
		for i := 0; i < size; i++ {
			start = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			*result = append(*result, start)

			// Push connected nodes to the stack.
			if len(plan[start]) > 0 {
				dest := plan[start]
				// Get all the destinations name.
				keys := make([]string, 0)
				for k, v := range dest {
					if v > 0 {
						keys = append(keys, k)
					}
				}
				// Sorting keys.
				sort.Strings(keys)
				// Then do traversal.
				for i := len(keys) - 1; i >= 0; i-- {
					stack = append(stack, keys[i])
					plan[start][keys[i]]--
				}
			}
		}
	}
}
