package week05onclass

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
// leetcode 153. 寻找旋转排序数组中的最小值

func findMin(nums []int) int {
	var left, right = 0, len(nums) - 1
	for left < right {
		var mid = (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
