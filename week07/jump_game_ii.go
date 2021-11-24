package week07

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/jump-game-ii/
// leetcode 45. 跳跃游戏 II

func jump(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	for i := 0; i < n ; i++ {
		f[i] = 1e9
	}
	f[0] = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= i + nums[i]; j++ {
			if j >= n {
				break
			}
			f[j] = internel.Min(f[j], f[i] + 1)
		}
	}
	return f[n - 1]
}
