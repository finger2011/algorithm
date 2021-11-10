package week06onclasss

// https://leetcode-cn.com/problems/maximum-product-subarray/
// leetcode 152. 乘积最大子数组

func maxProduct(nums []int) int {
	dp := make([][2]int, len(nums))
	dp[0] = [2]int{nums[0], nums[0]}
	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(nums[i], max(nums[i] * dp[i - 1][0],nums[i] * dp[i - 1][1]))
		dp[i][1] = min(nums[i], min(nums[i] * dp[i - 1][0],nums[i] * dp[i - 1][1]))
	}
	var ans = dp[0][0]
	for i := 1; i < len(dp); i++ {
		if dp[i][0] > ans {
			ans = dp[i][0]
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}