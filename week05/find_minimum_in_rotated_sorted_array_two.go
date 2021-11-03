package week05

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
// leetcode 寻找旋转排序数组中的最小值 II

func findMin(nums []int) int {
	var left, right = 0, len(nums) - 1
	for left < right {
		var mid = (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		}  else if nums[mid] > nums[right] {
			left = mid + 1 
		} else {
			right--
		}
	}
	return nums[right]
}