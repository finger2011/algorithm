package week04onclass

// https://leetcode-cn.com/problems/number-of-islands/
// 200. 岛屿数量

var visited map[int]map[int]bool
var gridData [][]byte
var dx, dy = []int{-1, 0, 0, 1}, []int{0, -1, 1, 0}

//DFS
func numIslandsByDFS(grid [][]byte) int {
	var ans int
	gridData = grid
	visited = make(map[int]map[int]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make(map[int]bool, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if string(grid[i][j]) == "1" && !visited[i][j] {
				ans++
				dfsIsland(i, j)
			}
		}
	}
	return ans
}

// BFS
func numIslandsByBFS(grid [][]byte) int {
	var ans int
	gridData = grid
	visited = make(map[int]map[int]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make(map[int]bool, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if string(grid[i][j]) == "1" && !visited[i][j] {
				ans++
				bfsIsland(i, j)
			}
		}
	}
	return ans
}

func dfsIsland(i, j int) {
	if i < 0 || i >= len(gridData) || j < 0 || j >= len(gridData[0]) {
		return
	}
	if string(gridData[i][j]) != "1" || visited[i][j] {
		return
	}

	visited[i][j] = true
	for k := 0; k < 4; k++ {
		dfsIsland(i+dx[k], j+dy[k])
	}
}

func bfsIsland(i, j int) {
	var queue = [][]int{{i, j}}
	visited[i][j] = true
	for len(queue) > 0 {
		var x, y = queue[0][0], queue[0][1]
		queue = queue[1:]
		for k := 0; k < 4; k++ {
			var nx, ny = x + dx[k], y + dy[k]
			if nx < 0 || nx >= len(gridData) || ny < 0 || ny >= len(gridData[0]) {
				continue
			}
			if string(gridData[nx][ny]) != "1" || visited[nx][ny] {
				continue
			}
			queue = append(queue, []int{nx, ny})
			visited[nx][ny] = true
		}
	}

}
