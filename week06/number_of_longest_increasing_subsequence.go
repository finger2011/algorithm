package week06

// https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/
// leetcode 673. 最长递增子序列的个数

func findNumberOfLIS(nums []int) int {
	dp := make([][2]int, len(nums))
	dp[0] = [2]int{1, 1}
	for i := 1; i < len(nums); i++ {
		dp[i] = [2]int{1, 1}
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i][0] < dp[j][0]+1 {
					dp[i][0] = dp[j][0] + 1
					dp[i][1] = dp[j][1]
				} else if dp[i][0] == dp[j][0]+1 {
					dp[i][1] += dp[j][1]
				}
			}
		}
	}
	var ans, max int
	for i := 0; i < len(dp); i++ {
		if dp[i][0] > max {
			ans = dp[i][1]
			max = dp[i][0]
		} else if dp[i][0] == max {
			ans += dp[i][1]
		}
	}
	return ans
}
