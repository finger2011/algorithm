package week07onclass

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
// leetcode 121. 买卖股票的最佳时机121. 买卖股票的最佳时机

func maxProfit(prices []int) int {
	var ans = 0
	var min = prices[0]
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if min < prices[i] {
			ans = maxP(ans, prices[i] - min)
		}
	}
	return ans
}

func maxP(a, b int) int {
	if a > b {
		return a
	}
	return b
}