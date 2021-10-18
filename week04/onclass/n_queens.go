package week04onclass

// https://leetcode-cn.com/problems/n-queens/
// 51. N 皇后
var num int
var usedCol map[int]bool
var usedPlus map[int]bool
var usedMinutes map[int]bool
var ansNQueen [][]int
var ansCol []int

func solveNQueens(n int) [][]string {
	num = n
	usedCol = make(map[int]bool, n)
	usedPlus = make(map[int]bool, n)
	usedMinutes = make(map[int]bool, n)
	ansNQueen = [][]int{}
	ansCol = []int{}
	dfsNQueen(0)
	var res = make([][]string, len(ansNQueen))
	for i := 0; i < len(ansNQueen); i++ {
		res[i] = make([]string, n)
		for j := 0; j < n; j++ {
			res[i][j] = ""
			for k := 0; k < n; k++ {
				if k == ansNQueen[i][j] {
					res[i][j] += "Q"
				} else {
					res[i][j] += "."
				}
				
			}
		}
	}
	return res
}

func dfsNQueen(row int) {
	if row == num {
		ansNQueen = append(ansNQueen, append([]int{}, ansCol...))
		return
	}
	for i := 0; i < num; i++ {
		if !usedCol[i] && !usedPlus[row+i] && !usedMinutes[row-i] {
			ansCol = append(ansCol, i)
			usedCol[i] = true
			usedPlus[row+i] = true
			usedMinutes[row-i] = true
			dfsNQueen(row + 1)
			usedMinutes[row-i] = false
			usedPlus[row+i] = false
			usedCol[i] = false
			ansCol = ansCol[0 : len(ansCol)-1]
		}
	}
}
