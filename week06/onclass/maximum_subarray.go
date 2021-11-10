package week06onclasss

// https://leetcode-cn.com/problems/maximum-subarray/
// leetcode 53. 最大子序和

func maxSubArray(nums []int) int {
    var ant = nums[0]
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {	
		if dp[i - 1] + nums[i] > nums[i] {
			dp[i] = dp[i - 1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if ant < dp[i] {
			ant = dp[i]
		}
	}
	return ant
}