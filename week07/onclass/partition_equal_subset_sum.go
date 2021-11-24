package week07onclass

// https://leetcode-cn.com/problems/partition-equal-subset-sum/
// leetcode 416. 分割等和子集

func canPartition(nums []int) bool {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum % 2 == 1 {
		return false
	}
	f := make(map[int]bool, sum/2 + 1)
	f[0] = true
	for i := 1; i <= len(nums); i++ {
		for j := sum/2; j >= nums[i - 1]; j-- {
			f[j] = f[j] || f[j - nums[i - 1]]
		}
	}
	return f[sum / 2]
}