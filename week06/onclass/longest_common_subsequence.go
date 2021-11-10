package week06onclasss

// https://leetcode-cn.com/problems/longest-common-subsequence/
// leetcode 1143. 最长公共子序列

func longestCommonSubsequence(text1 string, text2 string) int {
	text1 = " " + text1
	text2 = " " + text2
	grid := make([][]int, len(text1))
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, len(text2))
	}
	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			if text1[i] == text2[j] {
				grid[i][j] = grid[i-1][j-1] + 1
			} else {
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] = grid[i-1][j]
				} else {
					grid[i][j] = grid[i][j-1]
				}
			}
		}
	}
	return grid[len(text1)-1][len(text2)-1]
}
