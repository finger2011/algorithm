package week07

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/perfect-squares/
// leetcode 279. 完全平方数

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = 1e9
	}
	f[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			f[i] = internel.Min(f[i], f[i - j * j] + 1)
		}
	}
	return f[n]
}
