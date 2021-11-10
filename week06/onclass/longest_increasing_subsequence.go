package week06onclasss

// https://leetcode-cn.com/problems/longest-increasing-subsequence/
// leetcode 300. 最长递增子序列

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j] + 1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	var ans int
	for i := 0; i < len(dp); i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}