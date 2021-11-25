package week08onclass

// https://leetcode-cn.com/problems/number-of-provinces/
// leetcode 547. 省份数量
var fa []int

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	fa = make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				unionSet(i, j)
			}
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		if find(i) == i {
			ans++
		}
	}
	return ans
}

func find(x int) int {
	if x != fa[x] {
		fa[x] = find(fa[x])
	}
	return fa[x]
}

func unionSet(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		fa[x] = y
	}
}
