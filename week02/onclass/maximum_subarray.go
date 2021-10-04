package week02onclass

// https://leetcode-cn.com/problems/maximum-subarray/
// leetcode 53 最大子序和

func maxSubArray(nums []int) int {
	var ant = -100000
	var subMin int
	for i, pre := 0, 0; i < len(nums); i++ {
		pre += nums[i]
		if pre-subMin > ant {
			ant = pre - subMin
		}
		if pre < subMin {
			subMin = pre
		}
	}
	return ant
}

//前缀和
func maxSubArray2(nums []int) int {
	var ant = -100000
	var s = make([]int, len(nums)+1)
	s[0] = 0
	var preMin = make([]int, len(nums)+1)
	preMin[0] = s[0]
	for i := 1; i <= len(nums); i++ {
		s[i] = s[i-1] + nums[i-1]
		preMin[i] = s[i]
		if preMin[i-1] < s[i] {
			preMin[i] = preMin[i-1]
		}
		if s[i]-preMin[i-1] > ant {
			ant = s[i] - preMin[i-1]
		}
	}
	return ant
}

//动态规划
// f(n) = max(f(n -1) + nums[n], nums[n])
//原地存储
func maxSubArray3(nums []int) int {
	var ant = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i - 1] > nums[i] {
			nums[i] += nums[i - 1]
		}
		if nums[i] > ant {
			ant = nums[i]
		}
	}
	return ant
}
