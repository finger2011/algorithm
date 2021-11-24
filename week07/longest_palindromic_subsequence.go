package week07

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/longest-palindromic-subsequence/
// leetcode 516. 最长回文子序列

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j > i {
				dp[i][j] = 0
			} else if i == j {
				dp[i][j] = 1
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = internel.Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
