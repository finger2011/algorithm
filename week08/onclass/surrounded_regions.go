package week08onclass

// https://leetcode-cn.com/problems/surrounded-regions/
// leetcode 130. 被围绕的区域

// var dx, dy = []int{-1, 0, 0, 1}, []int{0, -1, 1, 0}
var regionFa []int

func solve(board [][]byte) [][]byte {
	m, n := len(board), len(board[0])
	regionFa = make([]int, m*n+1)
	for i := 0; i <= m*n; i++ {
		regionFa[i] = i
	}
	outside := m * n
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := i+dx[k], j+dy[k]
				if ni < 0 || nj < 0 || ni >= m || nj >= n {
					regionUnionSet(regionNum(i, j, n), outside)
				} else {
					if board[ni][nj] == 'O' {
						regionUnionSet(regionNum(i, j, n), regionNum(ni, nj, n))
					}
				}
			}
		}
	}
	outside = regionFind(outside)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && regionFind(regionNum(i, j, n)) != outside {
				board[i][j] = 'X'
			}
		}
	}
	return board
}

func regionNum(i, j, n int) int {
	return i*n + j
}

func regionUnionSet(x, y int) {
	x, y = regionFind(x), regionFind(y)
	if x != y {
		regionFa[x] = y
	}
}

func regionFind(x int) int {
	if regionFa[x] != x {
		regionFa[x] = regionFind(regionFa[x])
	}
	return regionFa[x]
}
