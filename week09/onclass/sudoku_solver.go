package week09onclass

import (
	"strconv"
)

// https://leetcode-cn.com/problems/sudoku-solver/
// leetcode 37. 解数独

// false未填，true已填
var row [][]bool
var col [][]bool
var box [][]bool

func solveSudoku(board [][]byte) {
	// 初始化为false - 未填
	row = make([][]bool, 9)
	col = make([][]bool, 9)
	box = make([][]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 9)
		col[i] = make([]bool, 9)
		box[i] = make([]bool, 9)
	}
	// 根据已有数字，设置为true - 已填
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			dig := int(board[i][j] - '1')
			row[i][dig] = true
			col[j][dig] = true
			box[boxIndex(i, j)][dig] = true
		}
	}
	sudokuDFS(board)
}

func sudokuDFS(board [][]byte) bool {
	// 获取可填数最少的位置，以及可填入的数
	pos, nums := findPos(board)
	if pos[0] == -1 {
		return true
	}
	for _, d := range nums {
		board[pos[0]][pos[1]] = strconv.Itoa(d + 1)[0]
		row[pos[0]][d] = true
		col[pos[1]][d] = true
		box[boxIndex(pos[0], pos[1])][d] = true
		if sudokuDFS(board) {
			return true
		}
		//还原现场
		row[pos[0]][d] = false
		col[pos[1]][d] = false
		box[boxIndex(pos[0], pos[1])][d] = false
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
			for d := 0; d < 9; d++ {
				// false为未填，所以需要取反
				if (!row[i][d]) && (!col[j][d]) && (!box[boxIndex(i, j)][d]) {
					c = append(c, d)
				}
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
