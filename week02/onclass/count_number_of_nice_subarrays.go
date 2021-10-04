package week02onclass

// https://leetcode-cn.com/problems/count-number-of-nice-subarrays/
// leetcode 1248 统计“优美子数组”

func numberOfSubarrays(nums []int, k int) int {
	var sum = make(map[int]int, len(nums)+1)
	var ant int
	sum[0] = 1
	for i, pre := 0, 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			pre++
		}
		if pre-k >= 0 {
			ant += sum[pre-k]
		}
		sum[pre]++
	}
	return ant
}
