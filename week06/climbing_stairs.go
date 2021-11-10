package week06

// https://leetcode-cn.com/problems/climbing-stairs/description/
// leetcode 70. 爬楼梯

var sMap = map[int]int{}

// 记忆化搜索
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	if _, ok := sMap[n]; ok {
		return sMap[n]
	}
	var ans = climbStairs(n-1) + climbStairs(n-2)
	sMap[n] = ans
	return ans
}
