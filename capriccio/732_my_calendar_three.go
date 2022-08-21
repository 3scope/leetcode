package main

import "sort"

// 类似于主持人问题，题目可以更改为随着不同安排的增加，最多需要多少主持人。
type MyCalendarThree struct {
	// 保存某个时刻最多需要多少人。
	Count map[int]int
	// 保存“map”中的最大值。
	Max int
}

func MyCalendarThreeConstructor() MyCalendarThree {
	return MyCalendarThree{
		Count: make(map[int]int),
		Max:   0,
	}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	// 差分数组。
	this.Count[start]++
	this.Count[end]--
	// 存储所有的时刻。
	times := make([]int, 0)
	for key, _ := range this.Count {
		times = append(times, key)
	}
	// 将时刻排序。
	sort.Ints(times)

	// 求前缀和，得到“k”的最大值。
	k := 0
	sum := 0
	for i := 0; i < len(times); i++ {
		sum += this.Count[times[i]]
		if sum > k {
			k = sum
		}
	}
	return k
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
