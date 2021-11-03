package week05

// https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/
// leetcode 1011. 在 D 天内送达包裹的能力
var numWeights []int

func shipWithinDays(weights []int, days int) int {
	var left, right = weights[0], 0
	for i := 0; i < len(weights); i++ {
		right += weights[i]
		if weights[i] > left {
			left = weights[i]
		}
	}
	numWeights = weights
	for left < right {
		var mid = (left + right) / 2
		if validDays(mid, days) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func validDays(cap, days int) bool {
	var day, everyDay = 1, 0
	for i := 0; i < len(numWeights); i++ {
		everyDay += numWeights[i]
		if everyDay > cap {
			day++
			everyDay = numWeights[i]
		}
	}
	return day <= days
}
