package week10onclass

// https://leetcode-cn.com/problems/power-of-two/
// leetcode 231. 2 的幂

func isPowerOfTwo(n int) bool {
	return n > 0 && n == (n & (-n))
}