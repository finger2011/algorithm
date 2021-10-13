package internel

import "reflect"

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
