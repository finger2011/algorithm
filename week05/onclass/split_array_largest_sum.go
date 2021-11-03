package week05onclass

// https://leetcode-cn.com/problems/split-array-largest-sum/
// leetcode 410. 分割数组的最大值
var numArr []int
func splitArray(nums []int, m int) int {
	var left, right = 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > left {
			left = nums[i]
		}
		right += nums[i]
	}
	numArr = nums
	for left < right {
		var mid = (left + right) / 2
		if validArray(mid, m) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func validArray(mid, m int) bool {
	var count = 1
	var every = 0
	for i := 0; i < len(numArr); i++ {
		every += numArr[i]
		if every > mid {
			every = numArr[i]
			count++
		}
	}
	return count <= m
}