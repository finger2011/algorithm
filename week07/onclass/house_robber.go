package week07onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/house-robber/
// leetcode 198. 打家劫舍

// 列表法
func rob1_2(nums []int) int {
	n := len(nums)
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = -1e9
		}
	}
	f[0][0] = 0
	for i := 0; i < n; i++ {
		f[i + 1][0] = internel.Max(f[i + 1][0], f[i][1])
		f[i + 1][0] = internel.Max(f[i + 1][0], f[i][0])
		f[i + 1][1] = f[i][0] + nums[i]
	}
	return internel.Max(f[n][0], f[n][1])
}

func rob1(nums []int) int {
	n := len(nums)
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = -1e9
		}
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		f[i][0] = internel.Max(f[i - 1][0], f[i - 1][1])
		f[i][1] = f[i - 1][0] + nums[i - 1]
	}
	return internel.Max(f[n][0], f[n][1])
}
