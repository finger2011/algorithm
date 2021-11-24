package week07onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
// leetcode 121. 买卖股票的最佳时机

func maxProfit(prices []int) int {
	var ans = 0
	var min = prices[0]
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if min < prices[i] {
			ans = internel.Max(ans, prices[i]-min)
		}
	}
	return ans
}
