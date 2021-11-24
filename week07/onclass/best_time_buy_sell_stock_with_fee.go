package week07onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
// leetcode 714. 买卖股票的最佳时机含手续费

func maxProfit5(prices []int, fee int) int {
	n := len(prices)
	prices = append([]int{0}, prices...)
	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = -1e9
		}
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		f[i][1] = internel.Max(f[i][1], f[i-1][0]-prices[i] - fee)
		f[i][0] = internel.Max(f[i][0], f[i-1][1]+prices[i])
		for j := 0; j < 2; j++ {
			f[i][j] = internel.Max(f[i][j], f[i-1][j])
		}

	}
	return f[n][0]
}
