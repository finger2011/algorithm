package week05

// https://leetcode-cn.com/problems/count-of-range-sum/
// leetcode 327. 区间和的个数
var sumLower, sumUpper int64

func countRangeSum(nums []int, lower int, upper int) int {
	sumLower, sumUpper = int64(lower), int64(upper)
	var prefix = make([]int64, len(nums)+1)
	prefix[0] = 0
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = int64(nums[i]) + prefix[i]
	}

	return mergeSort(prefix)
}

func mergeSort(nums []int64) int {
	if len(nums) <= 1 {
		return 0
	}
	var left = append([]int64{}, nums[0:len(nums)/2]...)
	var right = append([]int64{}, nums[len(nums)/2:]...)
	var ans = mergeSort(left) + mergeSort(right)
	var l, r int
	for i := 0; i < len(left); i++ {
		for l < len(right) && right[l]-left[i] < sumLower {
			l++
		}
		for r < len(right) && right[r]-left[i] <= sumUpper {
			r++
		}
		ans += r - l
	}
	l, r = 0, 0
	for i := 0; i < len(nums); i++ {
		if l < len(left) && (r >= len(right) || left[l] <= right[r]) {
			nums[i] = left[l]
			l++
		} else {
			nums[i] = right[r]
			r++
		}
	}
	return ans
}
