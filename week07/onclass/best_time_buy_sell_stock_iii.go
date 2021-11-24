package week07onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
// leetcode 123. 买卖股票的最佳时机 III

func maxProfit3(prices []int) int {
	k := 2
	n := len(prices)
	prices = append([]int{0}, prices...)
	f := make([][][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = make([]int, k+1)
			for m := 0; m <= k; m++ {
				f[i][j][m] = -1e9
			}
		}
	}
	f[0][0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j < 2; j++ {
			for m := 0; m <= k; m++ {
				if m > 0 {
					f[i][1][m] = internel.Max(f[i][1][m], f[i-1][0][m-1]-prices[i])
				}
				f[i][0][m] = internel.Max(f[i][0][m], f[i-1][1][m]+prices[i])
				f[i][j][m] = internel.Max(f[i][j][m], f[i-1][j][m])
			}
		}
	}
	var ans int
	ans = -1e9
	for i := 0; i <= k; i++ {
		ans = internel.Max(ans, f[n][0][i])
	}
	return ans
}
