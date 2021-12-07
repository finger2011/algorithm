package week09onclass

// https://leetcode-cn.com/problems/valid-sudoku/
// leetcode 36. 有效的数独

func isValidSudoku(board [][]byte) bool {
	rowMap := make([]map[byte]bool, 9)
	colMap := make([]map[byte]bool, 9)
	blockMap := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		rowMap[i] = make(map[byte]bool, 9)
		colMap[i] = make(map[byte]bool, 9)
		blockMap[i] = make(map[byte]bool, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := rowMap[i][board[i][j]]; ok {
				return false
			}
			rowMap[i][board[i][j]] = true
			if _, ok := colMap[j][board[i][j]]; ok {
				return false
			}
			colMap[j][board[i][j]] = true
			k := (i/3) * 3 + j/3
			if _, ok := blockMap[k][board[i][j]]; ok {
				return false
			}
			blockMap[k][board[i][j]] = true		
		}
	}
	return true
}