package week06onclasss

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
// leetcode 122. 买卖股票的最佳时机 II

func maxProfit(prices []int) int {
	var ans int
	for i := 1; i < len(prices); i++ {
		if prices[i] - prices[i - 1] > 0 {
			ans += prices[i] - prices[i - 1]
		}
	}
	return ans
}