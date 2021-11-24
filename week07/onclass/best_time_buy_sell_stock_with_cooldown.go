package week07onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// leetcode 309. 最佳买卖股票时机含冷冻期

// 滚动数组
func maxProfit6_3(prices []int) int {
	n := len(prices)
	prices = append([]int{0}, prices...)
	f := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		f[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = make([]int, 2)
			for l := 0; l < 2; l++ {
				f[i][j][l] = -1e9
			}
		}
	}
	f[0][0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			for l := 0; l < 2; l++ {
				if f[i&1][j][l] == -1e9 {
					continue
				}
				if j == 0 && l == 0 {
					f[(i+1)&1][1][0] = internel.Max(f[(i+1)&1][1][0], f[i&1][j][l]-prices[i+1])
				}
				if j == 1 && l == 0 {
					f[(i+1)&1][0][1] = internel.Max(f[(i+1)&1][0][1], f[i&1][j][l]+prices[i+1])
				}
				f[(i+1)&1][j][0] = internel.Max(f[(i+1)&1][j][0], f[i&1][j][l])
				f[i&1][j][l] = -1e9
			}
		}
	}
	return internel.Max(f[n&1][0][0], f[n&1][0][1])
}

//列表法
func maxProfit6_2(prices []int) int {
	n := len(prices)
	prices = append([]int{0}, prices...)
	f := make([][][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = make([]int, 2)
			for l := 0; l < 2; l++ {
				f[i][j][l] = -1e9
			}
		}
	}
	f[0][0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			for l := 0; l < 2; l++ {
				if f[i][j][l] == -1e9 {
					continue
				}
				if j == 0 && l == 0 {
					f[i+1][1][0] = internel.Max(f[i+1][1][0], f[i][j][l]-prices[i+1])
				}
				if j == 1 && l == 0 {
					f[i+1][0][1] = internel.Max(f[i+1][0][1], f[i][j][l]+prices[i+1])
				}
				f[i+1][j][0] = internel.Max(f[i+1][j][0], f[i][j][l])
			}
		}
	}
	return internel.Max(f[n][0][0], f[n][0][1])
}

func maxProfit6(prices []int) int {
	n := len(prices)
	prices = append([]int{0}, prices...)
	f := make([][][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = make([]int, 2)
			for l := 0; l < 2; l++ {
				f[i][j][l] = -1e9
			}
		}
	}
	f[0][0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j < 2; j++ {
			for l := 0; l < 2; l++ {
				f[i][1][0] = internel.Max(f[i][1][0], f[i-1][0][0]-prices[i])
				f[i][0][1] = internel.Max(f[i][0][1], f[i-1][1][0]+prices[i])
				f[i][j][0] = internel.Max(f[i][j][0], f[i-1][j][l])
			}
		}
	}
	return internel.Max(f[n][0][0], f[n][0][1])
}
