package week07

// https://leetcode-cn.com/problems/jump-game/
// leetcode 55. 跳跃游戏

func canJump(nums []int) bool {
	n := len(nums)
	f := make([]bool, n + 1)
	f[1] = true
	for i := 1; i <= n; i++ {
		if f[i] {
			for j := i; j <= i + nums[i - 1]; j++ {
				if j > n {
					break
				}
				f[j] = true
			}
		}
	}
	return f[n]
}