package week05onclass

// https://leetcode-cn.com/problems/find-peak-element/
// leetcode 162. 寻找峰值

func findPeakElement(nums []int) int {
	var left, right = 0, len(nums) - 1
	for left < right {
		var lmid = (left + right) / 2
		var rmid = lmid + 1
		if nums[lmid] <= nums[rmid] {
			left = lmid + 1
		} else {
			right = rmid - 1
		}
	}
	return right
}