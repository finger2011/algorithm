package week05onclass

// https://leetcode-cn.com/problems/binary-search/
// leetcode 704. 二分查找

func search(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	for left <= right {
		var mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}