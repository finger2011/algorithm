package week09onclass

// https://leetcode-cn.com/problems/distinct-subsequences/
// leetcode 115. 不同的子序列

func numDistinct(s string, t string) int {
	n, m := len(s), len(t)
	s = " " + s
	t = " " + t
	f := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m + 1)
		f[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] += f[i - 1][j]
			if s[i] == t[j] {
				f[i][j] += f[i - 1][j - 1]
			}
		}
	}
	return f[n][m]
}