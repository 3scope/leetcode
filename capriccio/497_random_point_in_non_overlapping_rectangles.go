package main

import (
	"math/rand"
	"sort"
)

type RectangleSolution struct {
	Rects [][]int
	Areas []int
}

func RectangleConstructor(rects [][]int) RectangleSolution {
	// 例如有三个矩形，每个矩形的面积依次为9，12，16，那么它们对应被选择的概率应为 9/37，12/37，16/37。那么随机选中“0~36”中间的一个点，如果是“0～8”，那么就相当于选中了第一个；如果是“9～20”，那么就相当于选中了第二个；如果是“21～36”，那么就相当于选中了第三个。
	areaPick := []int{0}
	// 构建前缀和，得到总面积。
	for i := 0; i < len(rects); i++ {
		area := (rects[i][2]-rects[i][0]+1)*(rects[i][3]-rects[i][1]+1) + areaPick[i]
		areaPick = append(areaPick, area)
	}

	return RectangleSolution{
		Rects: rects,
		Areas: areaPick,
	}
}

// 得到所有矩形的面积，根据面积来得到对应矩形被选择的概率。
// 之后从矩形中随机选取一个点。
func (rs *RectangleSolution) Pick() []int {
	rects := rs.Rects
	areas := rs.Areas
	// rand.Seed(time.Now().UnixNano())
	picked := rand.Intn(areas[len(areas)-1])
	// 二分查找，求的随机的点所在区间。
	index := sort.SearchInts(areas, picked+1) - 1
	// “middle”的值就是选中的第几个矩形，下标从0开始。
	pickedRect := rects[index]
	// 整数点的个数，需要加1。
	x := pickedRect[2] - pickedRect[0] + 1
	y := pickedRect[3] - pickedRect[1] + 1
	xOffset := rand.Intn(x)
	yOffset := rand.Intn(y)
	return []int{pickedRect[0] + xOffset, pickedRect[1] + yOffset}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
