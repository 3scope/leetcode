package main

type BitTrie struct {
	IsEnd bool
	Next  map[int]*BitTrie
}

func (bt *BitTrie) Insert(num int) {
	// The "num" is between 0 and 2^31 - 1, it has 31 bit length most, and can move 30 bit to left.
	moveToLeft := 30
	cur := bt
	for i := moveToLeft; i >= 0; i-- {
		bit := num >> i & 1
		if cur.Next[bit] == nil {
			temp := &BitTrie{
				Next: make(map[int]*BitTrie),
			}
			cur.Next[bit] = temp
			cur = temp
			continue
		} else {
			cur = cur.Next[bit]
		}
	}
	cur.IsEnd = true
}

func (bt *BitTrie) GetMaxXOROfNum(num int) int {
	moveToLeft := 30
	cur := bt
	result := 0
	for i := moveToLeft; i >= 0; i-- {
		bit := num >> i & 1
		if cur.Next[bit^1] == nil {
			cur = cur.Next[bit]
			result = result * 2
		} else {
			cur = cur.Next[bit^1]
			result = result*2 + 1
		}
	}
	return result
}

func findMaximumXOR(nums []int) int {
	root := &BitTrie{
		Next: make(map[int]*BitTrie),
	}
	result := 0
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for i := 1; i < len(nums); i++ {
		root.Insert(nums[i-1])
		result = max(result, root.GetMaxXOROfNum(nums[i]))
	}

	return result
}
