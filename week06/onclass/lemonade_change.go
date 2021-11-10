package week06onclasss

// https://leetcode-cn.com/problems/lemonade-change
// leetcode 860. 柠檬水找零
var hands map[int]int

func lemonadeChange(bills []int) bool {
	hands = make(map[int]int, 3)
	hands[5] = 0
	hands[10] = 0
	hands[20] = 0
	for i := 0; i < len(bills); i++ {
		hands[bills[i]]++
		if !exchange(bills[i] - 5) {
			return false
		}
	}
	return true
}

func exchange(amount int) bool {
	var coins = [3]int{20, 10, 5}
	for i := 0; i < len(coins); i++ {
		for amount >= coins[i] && hands[coins[i]] > 0 {
			amount -= coins[i]
			hands[coins[i]]--
		}
	}
	return amount == 0
}
