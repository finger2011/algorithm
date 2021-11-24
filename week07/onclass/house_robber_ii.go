package week07onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/house-robber-ii/
// leetcode 213. 打家劫舍 II

func rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = -1e9
		}
	}
	f[1][0] = 0
	for i := 2; i <= n; i++ {
		f[i][0] = internel.Max(f[i-1][0], f[i-1][1])
		f[i][1] = f[i-1][0] + nums[i-1]
	}
	ans := internel.Max(f[n][0], f[n][1])

	f[1][1] = nums[0]
	f[1][0] = 0
	for i := 2; i <= n; i++ {
		f[i][0] = internel.Max(f[i-1][0], f[i-1][1])
		f[i][1] = f[i-1][0] + nums[i-1]
	}
	return internel.Max(ans, f[n][0 ])
}
