package week10onclass

// https://leetcode-cn.com/problems/n-queens/
// leetcode 51. N 皇后

var num int
var usedCol int
var usedPlus int
var usedMinutes int
var ansNQueen [][]int
var ansCol []int
// var powers map[int]int

func solveNQueens(n int) [][]string {
	num = n
	usedCol, usedPlus, usedMinutes = 0, 0, 0
	ansNQueen = [][]int{}
	ansCol = []int{}
	powers = make(map[int]int, n)
	for i := 0; i < n; i++ {
		powers[1 << i] = i
	}
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
	state := (1 << num - 1) - 1 - (usedCol | usedPlus | usedMinutes)
	p, m := usedPlus, usedMinutes
	for state > 0 {
		lowbit := state & -state
		col := powers[lowbit]
		state -= lowbit
		ansCol = append(ansCol, col)
		usedCol = usedCol | (1 << col)
		usedPlus = usedPlus | (1 << col)
		usedMinutes = usedMinutes | (1 << col)
		usedPlus = usedPlus >> 1
		usedMinutes = (usedMinutes << 1) & ((1 << num) - 1)
		dfsNQueen(row + 1)
		usedMinutes = m
		usedPlus = p
		usedCol = usedCol ^ (1 << col)
		ansCol = ansCol[0 : len(ansCol)-1]
	}
	// for i := 0; i < num; i++ {
	// 	if !usedCol[i] && !usedPlus[row+i] && !usedMinutes[row-i] {
	// 		ansCol = append(ansCol, i)
	// 		usedCol[i] = true
	// 		usedPlus[row+i] = true
	// 		usedMinutes[row-i] = true
	// 		dfsNQueen(row + 1)
	// 		usedMinutes[row-i] = false
	// 		usedPlus[row+i] = false
	// 		usedCol[i] = false
	// 		ansCol = ansCol[0 : len(ansCol)-1]
	// 	}
	// }
}
