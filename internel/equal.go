package internel

import "reflect"

//Max 最大值
func Max(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}

//Min 最小值
func Min(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < ans {
			ans = nums[i]
		}
	}
	return ans
}

//IntsEqual [][]int equal
func IntsEqual(nums1, nums2 [][]int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for _, num := range nums1 {
		var isContain bool
		for _, compareStr := range nums2 {
			if IntEqual(num, compareStr) {
				isContain = true
				break
			}
		}
		if !isContain {
			return false
		}
	}
	return true
}

//IntEqual []int equal
func IntEqual(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for _, num := range nums1 {
		var isContain bool
		for _, compareStr := range nums2 {
			if reflect.DeepEqual(num, compareStr) {
				isContain = true
				break
			}
		}
		if !isContain {
			return false
		}
	}
	return true
}
