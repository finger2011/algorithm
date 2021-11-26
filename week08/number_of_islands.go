package week08

// https://leetcode-cn.com/problems/number-of-islands/
// leetcode 200. 岛屿数量

var dx = []int{1, 0, 0, -1}
var dy = []int{0, 1, -1, 0}
var islandFa []int

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	islandFa = make([]int, m*n)
	for i := 0; i < m*n; i++ {
		islandFa[i] = i
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			for k := 0; k < 4; k++ {
				nx, ny := i+dx[k], j+dy[k]
				if nx < 0 || ny < 0 || nx >= m || ny >= n  {
					continue
				}
				if grid[nx][ny] == '1' {
					islandUnion(i*n+j, nx*n+ny)
				}
			}
		}
	}
	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && islandFind(i*n+j) == i*n+j {
				ans++
			}
		}
	}
	return ans
}

func islandFind(x int) int {
	if islandFa[x] != x {
		islandFa[x] = islandFind(islandFa[x])
	}
	return islandFa[x]
}

func islandUnion(x, y int) {
	x, y = islandFind(x), islandFind(y)
	if x != y {
		islandFa[x] = y
	}
}
