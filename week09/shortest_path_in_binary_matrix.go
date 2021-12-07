package week09

// https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/
// leetcode 1091. 二进制矩阵中的最短路径

var pathGrid [][]int

//双向BFS
func shortestPathBinaryMatrix2(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	n := len(grid)
	if n == 1 {
		return 1
	}
	if grid[n-1][n-1] == 1 {
		return -1
	}
	pathGrid = grid
	depBegin := make([][]int, n)
	depEnd := make([][]int, n)
	for i := 0; i < n; i++ {
		depBegin[i] = make([]int, n)
		depEnd[i] = make([]int, n)
		for j := 0; j < n; j++ {
			depBegin[i][j] = -1
			depEnd[i][j] = -1
		}
	}
	depBegin[0][0] = 0
	depEnd[n-1][n-1] = 0
	qBegin := [][]int{{0, 0}}
	qEnd := [][]int{{n - 1, n - 1}}
	var res int
	for len(qEnd) > 0 || len(qBegin) > 0 {
		res, qBegin = pathBFS(qBegin, depBegin, depEnd, n)
		if res != -1 {
			return res + 1
		}
		res, qEnd = pathBFS(qEnd, depEnd, depBegin, n)
		if res != -1 {
			return res + 1
		}
	}
	return -1
}

func pathBFS(q [][]int, dep1, dep2 [][]int, n int) (int, [][]int) {
	var dx = [8]int{0, 1, 1, 1, 0, -1, -1, -1}
	var dy = [8]int{-1, -1, 0, 1, 1, 1, 0, -1}
	newQ := make([][]int, 0)
	for _, pos := range q {
		for i := 0; i < 8; i++ {
			nx, ny := pos[0]+dx[i], pos[1]+dy[i]
			if nx < 0 || nx >= n || ny < 0 || ny >= n || pathGrid[nx][ny] == 1 || dep1[nx][ny] != -1 {
				continue
			}
			dep1[nx][ny] = dep1[pos[0]][pos[1]] + 1
			newQ = append(newQ, []int{nx, ny})
			if dep2[nx][ny] != -1 {
				return dep1[nx][ny] + dep2[nx][ny], q
			}
		}
	}
	return -1, newQ
}

// 普通DFS
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	dx := [8]int{0, 1, 1, 1, 0, -1, -1, -1}
	dy := [8]int{-1, -1, 0, 1, 1, 1, 0, -1}
	n := len(grid)
	q := [][]int{{0, 0}}
	dep := make([][]int, n)
	for i := 0; i < n; i++ {
		dep[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dep[i][j] = -1
		}
	}
	dep[0][0] = 1
	for len(q) > 0 {
		pos := q[0]
		q = q[1:]
		for i := 0; i < 8; i++ {
			nx, ny := pos[0]+dx[i], pos[1]+dy[i]
			if nx < 0 || nx >= n || ny < 0 || ny >= n || grid[nx][ny] == 1 || dep[nx][ny] > 0 {
				continue
			}
			dep[nx][ny] = dep[pos[0]][pos[1]] + 1
			q = append(q, []int{nx, ny})
		}
	}
	return dep[n-1][n-1]
}
