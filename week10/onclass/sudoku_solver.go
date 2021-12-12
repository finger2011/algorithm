package week10onclass

import "strconv"

// https://leetcode-cn.com/problems/sudoku-solver/
// leetcode 37. 解数独

var row []int
var col []int
var box []int
var powers map[int]int

func solveSudoku(board [][]byte) {
	powers = make(map[int]int, 9)
	for i := 1; i <= 9; i++ {
		powers[1 << i] = i
	}
	row = make([]int, 9)
	col = make([]int, 9)
	box = make([]int, 9)
	for i := 0; i < 9; i++ {
		row[i] = ((1 << 9) - 1)
		col[i] = ((1 << 9) - 1)
		box[i] = ((1 << 9) - 1)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			dig := int(board[i][j] - '1')
			if (row[i] >> dig) & 1 == 1 {
				row[i] = row[i] ^ (1 << dig)
			} 
			if (col[j] >> dig) & 1 == 1 {
				col[j] = col[j] ^ (1 << dig)
			} 
			if (box[boxIndex(i, j)] >> dig) & 1 == 1 {
				box[boxIndex(i, j)] = box[boxIndex(i, j)] ^ (1 << dig)
			} 
			// row[i][dig] = true
			// col[j][dig] = true
			// box[boxIndex(i, j)][dig] = true
		}
	}
	sudokuDFS(board)

}

func sudokuDFS(board [][]byte) bool {
	pos, nums := findPos(board)
	if pos[0] == -1 {
		return true
	}
	for _, d := range nums {
		board[pos[0]][pos[1]] = strconv.Itoa(d + 1)[0]
		row[pos[0]] = row[pos[0]] ^ (1 << d)
		col[pos[1]] = col[pos[1]] ^ (1 << d)
		box[boxIndex(pos[0], pos[1])] = box[boxIndex(pos[0], pos[1])] ^ (1 << d)
		// row[pos[0]][d] = true
		// col[pos[1]][d] = true
		// box[boxIndex(pos[0], pos[1])][d] = true
		if sudokuDFS(board) {
			return true
		}
		box[boxIndex(pos[0], pos[1])] = box[boxIndex(pos[0], pos[1])] | (1 << d)
		col[pos[1]] = col[pos[1]] | (1 << d)
		row[pos[0]] = row[pos[0]] | (1 << d)
		// row[pos[0]][d] = false
		// col[pos[1]][d] = false
		// box[boxIndex(pos[0], pos[1])][d] = false
		board[pos[0]][pos[1]] = '.'
	}
	return false
}

func findPos(board [][]byte) ([]int, []int) {
	size := 10
	pos := []int{-1, -1}
	nums := []int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			c := []int{}
			state := row[i] & col[j] & box[boxIndex(i, j)]
			// for d := 0; d < 9; d++ {
			// 	if (state >> d) & 1 == 1 {
			// 		c = append(c, d)
			// 	}
			// }
			for state > 0 {
                lowbit := state & -state
				c = append(c, powers[lowbit])
                state -= lowbit
			}
			if len(c) < size {
				size = len(c)
				pos = []int{i, j}
				nums = c
			}
		}
	}
	return pos, nums
}

func boxIndex(i, j int) int {
	return (i/3)*3 + (j / 3)
}
