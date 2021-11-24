package week07onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/edit-distance/
// leetcode 72. 编辑距离

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	word1 = " " + word1
	word2 = " " + word2
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			f[i][j] = 1e9
		}
	}
	for i := 0; i <= n; i++ {
		f[i][0] = i
	}
	for j := 0; j <= m; j++ {
		f[0][j] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			var eq int
			if word1[i] != word2[j] {
				eq = 1
			}
			f[i][j] = internel.Min(f[i][j - 1] + 1, f[i - 1][j] + 1, f[i - 1][j - 1] + eq)
		}
	}
	return f[n][m]
}
