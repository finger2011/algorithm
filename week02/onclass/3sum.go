package week02onclass

import (
	"sort"
)

// https://leetcode-cn.com/problems/3sum/
// leetcode 15. 三数之和

func threeSum(nums []int, target int) [][]int {
	//排序
	var ant = [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		//防止重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//对i+1~len(nums)-1进行双指针扫描
		for j, k := i+1, len(nums)-1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[i]+nums[j]+nums[k] > target {
				k--
			}
			if j < k && nums[i]+nums[j]+nums[k] == target {
				ant = append(ant, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ant
}
