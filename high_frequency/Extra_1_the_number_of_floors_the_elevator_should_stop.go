package main

// 电梯里面载有一定数量的乘客，现在从一楼出发，只停一次，停在哪一层所有乘客的总爬楼层数最少。
// 根据乘客想去的楼层，计算出电梯需要停留的层数。
// 电梯会尽量向地楼层停留。
func getMinFloor(layers map[int]int, height int) int {
	// 默认电梯停在一楼。
	result := 1
	// 需要爬楼的总数。
	floors := 0
	// 初始化，从第一层开始。
	// p1 代表有多少人目的地在当前层数之下。
	p1 := 0
	// p2 代表有多少人目的地就是当前层数。
	p2 := layers[1]
	// p3 代表有多少人目的地在当前层数之上。
	p3 := 0
	for k, v := range layers {
		if k > 1 {
			p3 += v
			floors += (k - 1) * v
		}
	}
	// 电梯每上一层，有 p1 + p2 的人需要多爬一层楼，p3 的人可以少爬一层楼。
	// 也就是说，如果 p1 + p2 >= p3 那么停到当前层即可。
	for i := 1; i < height; i++ {
		if p1+p2 >= p3 {
			break
		} else {
			// 就往上爬一层。
			result++
			floors -= p3 - p1 - p2
			// p1 / p2 / p3 需要更新。
			p1 += layers[i]
			p2 = layers[i+1]
			p3 -= layers[i+1]
		}
	}
	return result
}
