package week06onclasss

//https://leetcode-cn.com/problems/coin-change/
// leetcode 322. 零钱兑换

//动规 递推
func coinChange(coins []int, amount int) int {
	var opt = make([]int, amount+1)
	opt[0] = 0
	for i := 1; i <= amount; i++ {
		opt[i] = 1e9
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && opt[i] > opt[i-coins[j]]+1 {
				opt[i] = opt[i-coins[j]] + 1
			}
		}
	}
	if opt[amount] >= 1e9 {
		opt[amount] = -1
	}
	return opt[amount]
}

var tCoins []int
var coinOpt []int

//动规 记忆化搜索
func coinChange1(coins []int, amount int) int {
	tCoins = coins
	coinOpt = make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		coinOpt[i] = -1
	}
	ans := calc(amount)
	if ans >= 1e9 {
		return -1
	}
	return ans
}

func calc(amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return 1e9
	}
	if coinOpt[amount] != -1 {
		return coinOpt[amount]
	}
	coinOpt[amount] = 1e9
	for i := 0; i < len(tCoins); i++ {
		var ans = calc(amount-tCoins[i]) + 1
		if ans < coinOpt[amount] {
			coinOpt[amount] = ans
		}
	}
	return coinOpt[amount]
}
