package week07onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/
// leetcode 1000. 合并石头的最低成本

func mergeStones(stones []int, k int) int {
	n := len(stones)
	f := make([][][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = make([]int, k+1)
			for m := 0; m < k+1; m++ {
				f[i][j][m] = 1e9
			}
		}
	}
	for i := 0; i < n; i++ {
		f[i][i][1] = 0
	}
	for le := 2; le <= n; le++ {
		for l := 0; l < n-le+1; l++ {
			r := l + le - 1
			for i := 2; i <= k; i++ {
				for p := l; p < r; p++ {
					f[l][r][i] = internel.Min(f[l][r][i], f[l][p][1]+f[p+1][r][i-1])
				}
			}
			sum := 0
			for p := l; p <= r; p++ {
				sum += stones[p]
			}
			f[l][r][1] = internel.Min(f[l][r][1], f[l][r][k]+sum)
		}
	}
	if f[0][n-1][1] == 1e9 {
		return -1
	}
	return f[0][n-1][1]
}
