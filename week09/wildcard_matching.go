package week09

// https://leetcode-cn.com/problems/wildcard-matching/
// leetcode 44. 通配符匹配

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	s = " " + s
	p = " " + p
	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 1; i <= m; i++{
		if p[i] == '*' {
			f[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j] >= 'a' && p[j] <= 'z' {
				f[i][j] = f[i-1][j-1] && s[i] == p[j]
			} else if p[j] == '?' {
				f[i][j] = f[i-1][j-1]
			} else {
				if j == 1 {
					f[i][j] = true
				} else {
					f[i][j] = f[i][j - 1] || f[i - 1][j]
				}
			}
		}
	}
	return f[n][m]
}