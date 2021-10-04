package week02

// https://leetcode-cn.com/problems/subarray-sum-equals-k/
// leetcode 560. 和为 K 的子数组

func subarraySum(nums []int, k int) int {
	var ant int
	var numMap = make(map[int]int, len(nums)+1)
	numMap[0] = 1
	for i, pre := 0, 0; i < len(nums); i++ {
		pre += nums[i]
		ant += numMap[pre-k]
		numMap[pre]++
	}
	return ant
}
