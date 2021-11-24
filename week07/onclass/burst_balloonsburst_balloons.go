package week07onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/burst-balloons/
// leetcode 312. 戳气球

// 记忆化搜索
var coinNum []int
var f [][]int

func maxCoins(nums []int) int {
	n := len(nums)
	coinNum = append(append([]int{1}, nums...), 1)
	f = make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		f[i] = make([]int, n+2)
		for j := 0; j < n+2; j++ {
			f[i][j] = -1
		}
	}
	return coinCal(1, n)
}

func coinCal(l, r int) int {
	if l > r {
		return 0
	}
	if f[l][r] != -1 {
		return f[l][r]
	}
	for p := l; p <= r; p++ {
		f[l][r] = internel.Max(f[l][r], coinCal(l, p-1)+coinCal(p+1, r)+coinNum[p]*coinNum[l-1]*coinNum[r+1])
	}
	return f[l][r]
}

func maxCoins2(nums []int) int {
	t := len(nums)
	nums = append(append([]int{1}, nums...), 1)
	f := make([][]int, t+2)
	for i := 0; i < t+2; i++ {
		f[i] = make([]int, t+2)
	}
	for length := 1; length <= t; length++ {
		for l := 1; l <= t-length+1; l++ {
			r := l + length - 1
			for p := l; p <= r; p++ {
				f[l][r] = internel.Max(f[l][r], f[l][p-1]+f[p+1][r]+nums[p]*nums[l-1]*nums[r+1])
			}
		}
	}
	return f[1][t]
}
