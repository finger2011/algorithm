package week02onclass

import (
	"sort"
)

//https://leetcode-cn.com/problems/two-sum/
// leetcode 1 两数之和

//双指针扫描
func twoSum3(nums []int, target int) []int {
	//先排序
	var sortNums [][]int
	for i := 0; i < len(nums); i++ {
		sortNums = append(sortNums, []int{nums[i], i})
	}
	sort.Slice(sortNums, func(i, j int) bool {
		return sortNums[i][0] < sortNums[j][0]
	})
	//双指针扫描
	for i, j := 0, len(sortNums)-1; i < len(sortNums); i++ {
		for i < j && sortNums[i][0]+sortNums[j][0] > target {
			j--
		}
		if i < j && sortNums[i][0]+sortNums[j][0] == target {
			return []int{sortNums[i][1], sortNums[j][1]}
		}
	}
	return []int{}
}

//两次遍历
func twoSum(nums []int, target int) []int {
	var numMap = make(map[int][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = append(numMap[nums[i]], i)
	}
	for i := 0; i < len(nums); i++ {
		if len(numMap[target-nums[i]]) > 0 {
			if target == 2*nums[i] {
				if len(numMap[target-nums[i]]) > 1 {
					return []int{numMap[nums[i]][0], numMap[nums[i]][1]}
				}
			} else {
				return []int{i, numMap[target-nums[i]][0]}
			}
		}
	}
	return []int{-1, -1}
}

// 一次遍历
func twoSum2(nums []int, target int) []int {
	var numMap = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if offset, ok := numMap[target-nums[i]]; ok {
			return []int{offset, i}
		}
		numMap[nums[i]] = i
	}
	return []int{-1, -1}
}
