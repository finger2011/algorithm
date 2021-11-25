package week08

// https://leetcode-cn.com/problems/redundant-connection/
// leetcode 684. 冗余连接
var redundantFa []int

func findRedundantConnection(edges [][]int) []int {
	redundantFa = make([]int, len(edges) + 1)
	for i := 0; i < len(redundantFa); i++ {
		redundantFa[i] = i
	}
	for i := 0; i < len(edges); i++ {
		if redundantFind(edges[i][0]) == redundantFind(edges[i][1]) {
			return edges[i]
		}
		redundantUnion(edges[i][0], edges[i][1])
	}
	return []int{-1, -1}
}

func redundantFind(x int) int {
	if redundantFa[x] != x {
		redundantFa[x] = redundantFind(redundantFa[x])
	}
	return redundantFa[x]
}

func redundantUnion(x, y int) {
	x, y = redundantFind(x), redundantFind(y)
	if x != y {
		redundantFa[x] = y
	}
}
