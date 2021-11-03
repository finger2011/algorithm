package week05

// https://leetcode-cn.com/problems/koko-eating-bananas/
// leetcode 875. 爱吃香蕉的珂珂
var numPiles []int

func minEatingSpeed(piles []int, h int) int {
	var left, right = 1, 0
	for i := 0; i < len(piles); i++ {
		right += piles[i]
	}
	numPiles = piles
	for left < right {
		var mid = (left + right) / 2
		if validEating(mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func validEating(speed, h int) bool {
	var hour = 0
	for i := 0; i < len(numPiles); i++ {
		hour += numPiles[i] / speed
		if numPiles[i]%speed > 0 {
			hour++
		}
	}
	return hour <= h
}
