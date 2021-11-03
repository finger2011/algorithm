package week05onclass

// https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/
// leetcode 1482. 制作 m 束花所需的最少天数
var numBloomDay []int

func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) == 0 {
		return -1
	}
	var left, right = bloomDay[0], 0
	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] < left {
			left = bloomDay[i]
		}
		if bloomDay[i] > right {
			right = bloomDay[i]
		}
	}
	var max = right
	right++
	numBloomDay = bloomDay
	for left < right {
		var mid = (left + right) / 2
		if validBloomDay(mid, m, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if right == max + 1 {
		return -1
	}
	return right
}

func validBloomDay(now, m, k int) bool {
	var count, every = 0, 0
	for i := 0; i < len(numBloomDay); i++ {
		if numBloomDay[i] <= now {
			every++
			if every >= k {
				every = 0
				count++
			}
		} else {
			every = 0
		}
	}
	return count >= m
}
