package week07onclass

// https://leetcode-cn.com/problems/coin-change-2/
// leetcode 518. 零钱兑换 II

func change(amount int, coins []int) int {
	n := len(coins)
	f := make([]int, amount + 1)
	f[0] = 1
	for i := 0; i < n; i++ {
		for j := coins[i]; j <= amount; j++{
			f[j] += f[j - coins[i]]
		}
	}
	return f[amount]
}