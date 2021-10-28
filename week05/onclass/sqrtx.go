package week05onclass

// https://leetcode-cn.com/problems/sqrtx/
// leetcode 69. Sqrt(x)

func mySqrt(x int) int {
	var left, right = 0, x
	for left < right {
		var mid = (left + right + 1) / 2
		if mid * mid <= x {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}