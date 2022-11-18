package main

// “DFS”方式。
func findCircleNum(isConnected [][]int) int {
	// 使用一个数组标记哪些省份被访问到了。
	isVisited := make([]bool, len(isConnected))
	result := 0
	// 由于是非联通图，所以需要多次调用“DFS”。
	for i := 0; i < len(isConnected); i++ {
		// 如果没有访问到，那么从该省份为起点进行“DFS”遍历。
		if !isVisited[i] {
			result++
			dfsProvinces(isConnected, isVisited, i)
		}
	}
	return result
}

func dfsProvinces(isConnected [][]int, isVisited []bool, index int) {
	// 如果之前访问过，那么直接返回。
	if isVisited[index] {
		return
	}

	// 没有访问过，就先标记。
	isVisited[index] = true
	for i := 0; i < len(isConnected[index]); i++ {
		// 如果两个省之间是联通的话，那么就递归进行访问。
		if isConnected[index][i] == 1 {
			dfsProvinces(isConnected, isVisited, i)
		}
	}
}

func findCircleNumBFS(isConnected [][]int) int {
	// 使用一个数组标记哪些省份被访问到了。
	isVisited := make([]bool, len(isConnected))
	result := 0
	// 由于是非联通图，所以需要多次调用“DFS”。
	for i := 0; i < len(isConnected); i++ {
		if !isVisited[i] {
			result++

		}
	}

	return result
}

func bfsProvinces(isConnected [][]int, isVisited []bool, index int) {
	if isVisited[index] {
		return
	}

	// “BFS”需要使用队列进行辅助。
	queue := make([]int, 0)
}
