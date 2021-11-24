package week07onclass

import (
	"finger2011/algggorithm/internel"
)

// https://leetcode-cn.com/problems/maximum-sum-circular-subarray/
// leetcode 918. 环形子数组的最大和

	

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	var ans = int(-1e9)

	var pre = make([]int, 2 * n + 1)
	pre[0] = 0
	for i := 1; i <= n; i++ {
		pre[i] = pre[i - 1] + nums[i - 1]
	}
	for i := n + 1; i <= 2 * n ; i++ {
		pre[i] = pre[i - 1] + nums[i - n - 1]
	}
	var q []int
	for i := 1; i <= 2 * n; i++ {
		for len(q) > 0 && q[0] < i - n {
			q = q[1:]
		}
		if len(q) > 0 {
			ans = internel.Max(ans, pre[i] - pre[q[0]])
		}
		for len(q) > 0 &&  pre[q[len(q) - 1]] >= pre[i] {
			q = q[0:len(q) - 1]
		}
		q = append(q, i)
	}
	return ans
}

func maxSubarraySumCircular2(nums []int) int {
	n := len(nums)
	var ans = int(-1e9)
	// 不跨过n
	var pre = make([]int, n + 1)
	pre[0] = 0
	for i := 0; i < n; i++ {
		pre[i + 1] = pre[i] + nums[i]
	}
	var tmp = int(1e9)
	for i := 1; i <= n; i++ {
		tmp = internel.Min(tmp, pre[i - 1])
		ans = internel.Max(ans, pre[i] - tmp)
	}
	
	tmp = -1e9
	ansMin := int(1e9)
	for i := 2; i <= n; i++ {
		tmp = internel.Max(tmp, pre[i - 1])
		ansMin = internel.Min(ansMin, pre[i] - tmp)
	}
	for i := 1; i < n; i++ {
		ansMin = internel.Min(ansMin, pre[i])
	}
	return internel.Max(ans, pre[n] - ansMin)
}