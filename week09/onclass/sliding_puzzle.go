package week09onclass

import "strconv"

// https://leetcode-cn.com/problems/sliding-puzzle/
// leetcode 773. 滑动谜题

var q []string
var sDep map[string]int

func slidingPuzzle(board [][]int) int {
	start := zip(board)
	target := zip([][]int{{1, 2, 3}, {4, 5, 0}})
	q = []string{start}
	sDep = make(map[string]int, 1)
	sDep[start] = 0
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		pos := findZero(s)
		if pos >= 3 {
			sExpand(s, pos, pos-3)
		} else {
			sExpand(s, pos, pos+3)
		}
		if pos%3 != 0 {
			sExpand(s, pos, pos-1)
		}
		if pos%3 != 2 {
			sExpand(s, pos, pos+1)
		}
		if _, ok := sDep[target]; ok {
			return sDep[target]
		}
	}
	return -1
}

func sExpand(s string, p1, p2 int) {
	var b = []byte(s)
	tmp := b[p1]
	b[p1] = b[p2]
	b[p2] = tmp
	if _, ok := sDep[string(b)]; ok {
		return
	}
	sDep[string(b)] = sDep[s] + 1
	q = append(q, string(b))
}

func findZero(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			return i
		}
	}
	return -1
}

func zip(board [][]int) string {
	var ans string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			ans += strconv.Itoa(board[i][j])
		}
	}
	return ans
}
