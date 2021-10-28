package week05onclass

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// leetcode 34. 在排序数组中查找元素的第一个和最后一个位置

func searchRange(nums []int, target int) []int {
	var left, right = -1, len(nums) - 1
	for left < right {
		var mid = (left + right + 1) / 2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	var end = right
	left, right = 0, len(nums)
	for left < right {
		var mid = (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	var start = left
	if end < start {
		return []int{-1, -1}
	}
	return []int{start, end}
}
