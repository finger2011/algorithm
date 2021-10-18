package week04

// https://leetcode-cn.com/problems/surrounded-regions/
// 130. 被围绕的区域

// 从四条边遍历（BFS or DFS），所有的连通图都是不可修改
// 再从剩余遍历的连通图都是可以修改的
var boardData [][]byte
var dx, dy = []int{-1, 0, 0, 1}, []int{0, -1, 1, 0}
var visited map[int]map[int]bool

//BFS
func solveBFS(board [][]byte) [][]byte {
	boardData = board
	visited = make(map[int]map[int]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make(map[int]bool, len(board[i]))
	}
	//优先遍历四条边
	for i := 0; i < len(board); i++ {
		if i == 0 || i == len(board)-1 {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == 'O' && !visited[i][j] {
					bfsSolve(i, j, false)
				}
			}
		} else {
			if board[i][0] == 'O' && !visited[i][0] {
				bfsSolve(i, 0, false)
			}
			if board[i][len(board[i])-1] == 'O' && !visited[i][len(board[i])-1] {
				bfsSolve(i, len(board[i])-1, false)
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 1; j < len(board[i])-1; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				bfsSolve(i, j, true)
			}
		}
	}
	board = boardData
	return boardData
}

//DFS
func solveDFS(board [][]byte) [][]byte {
	boardData = board
	visited = make(map[int]map[int]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make(map[int]bool, len(board[i]))
	}
	//优先遍历四条边
	for i := 0; i < len(board); i++ {
		if i == 0 || i == len(board)-1 {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == 'O' && !visited[i][j] {
					dfsSolve(i, j, -1, -1, false)
				}
			}
		} else {
			if board[i][0] == 'O' && !visited[i][0] {
				dfsSolve(i, 0, -1, -1, false)
			}
			if board[i][len(board[i])-1] == 'O' && !visited[i][len(board[i])-1] {
				dfsSolve(i, len(board)-1, -1, -1, false)
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 1; j < len(board[i])-1; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				dfsSolve(i, j, -1, -1, true)
			}
		}
	}
	board = boardData
	return boardData
}

func bfsSolve(i, j int, change bool) {
	var queue = [][]int{{i, j}}
	for len(queue) > 0 {
		var x, y = queue[0][0], queue[0][1]
		queue = queue[1:]
		visited[x][y] = true
		if change {
			boardData[x][y] = 'X'
		}
		for k := 0; k < 4; k++ {
			var nx, ny = x + dx[k], y + dy[k]
			if nx < 0 || nx >= len(boardData) || ny < 0 || ny >= len(boardData[0]) {
				continue
			}
			if !visited[nx][ny] && boardData[nx][ny] == 'O' {
				bfsSolve(nx, ny, change)
			}
		}
	}
}

func dfsSolve(i, j int, fatheri, fatherj int, change bool) {
	if i < 0 || i >= len(boardData) || j < 0 || j >= len(boardData[0]) || boardData[i][j] == 'X' {
		return
	}
	visited[i][j] = true
	for k := 0; k < 4; k++ {
		var x, y = i + dx[k], j + dy[k]
		if x == fatheri && y == fatherj {
			continue
		}
		if !visited[x][y] {
			dfsSolve(x, y, i, j, change)
		}
	}
	if change {
		boardData[i][j] = 'X'
	}
}
