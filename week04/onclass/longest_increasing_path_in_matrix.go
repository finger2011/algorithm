package week04onclass

import (
	"finger2011/algggorithm/internel"
)

// https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/
// leetcode 329. 矩阵中的最长递增路径

// var dx, dy = []int{-1, 0, 0, 1}, []int{0, -1, 1, 0}

//DFS
var matrixData [][]int
var dep [][]int
var ansPath int

func longestIncreasingPath2(matrix [][]int) int {
	matrixData = matrix
	dep = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dep[i] = make([]int, len(matrix[i]))
	}
	ansPath = 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			dfsPath(i, j)
		}
	}
	return ansPath
}

func dfsPath(i, j int) int {
	if dep[i][j] > 0 {
		return dep[i][j]
	}
	dep[i][j] = 1
	for k := 0; k < 4; k++ {
		var nx, ny = i + dx[k], j + dy[k]
		if nx < 0 || ny < 0 || nx >= len(matrixData) || ny >= len(matrixData[0]) {
			continue
		}
		if matrixData[nx][ny] > matrixData[i][j] {
			dep[i][j] = internel.Max(dep[i][j], dfsPath(nx, ny)+1)
		}
	}
	ansPath = internel.Max(ansPath, dep[i][j])
	return dep[i][j]
}

//BFS
func longestIncreasingPath(matrix [][]int) int {
	var to = make(map[int][]int)
	var m, n = len(matrix), len(matrix[0])
	var deg, dist = make([]int, m*n), make([]int, m*n)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for k := 0; k < 4; k++ {
				var nx, ny = i + dx[k], j + dy[k]
				if nx < 0 || ny < 0 || nx >= len(matrix) || ny >= len(matrix[0]) {
					continue
				}
				if matrix[nx][ny] > matrix[i][j] {
					if len(to[i*n+j]) == 0 {
						to[i*n+j] = []int{nx*n + ny}
					} else {
						to[i*n+j] = append(to[i*n+j], nx*n+ny)
					}
					deg[nx*n+ny]++
				}
			}
		}
	}

	var queue []int
	for i := 0; i < len(deg); i++ {
		if deg[i] == 0 {
			queue = append(queue, i)
			dist[i] = 1
		}
	}
	var ans int
	for len(queue) > 0 {
		var n = queue[0]
		queue = queue[1:]
		for i := 0; i < len(to[n]); i++ {
			deg[to[n][i]]--
			dist[to[n][i]] = internel.Max(dist[n]+1, dist[to[n][i]])
			if deg[to[n][i]] == 0 {
				queue = append(queue, to[n][i])
			}
		}
	}

	for i := 0; i < m*n; i++ {
		ans = internel.Max(ans, dist[i])
	}
	return ans
}
