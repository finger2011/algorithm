package week06

// https://leetcode-cn.com/problems/triangle
// leetcode 120. 三角形最小路径和

// 时间:O(n*n) 空间:O(n*n)
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, len(triangle))
	dp[len(dp)-1] = triangle[len(triangle)-1]
	for i := len(triangle) - 1; i > 0; i-- {
		dp[i-1] = make([]int, len(triangle[i-1]))
		for j := 1; j < len(triangle[i]); j++ {
			dp[i-1][j-1] = min(dp[i][j-1], dp[i][j]) + triangle[i-1][j-1]
		}
	}
	return dp[0][0]
}
// 时间:O(n*n) 空间:O(n)
func minimumTotal2(triangle [][]int) int {
	dp := triangle[len(triangle)-1]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
