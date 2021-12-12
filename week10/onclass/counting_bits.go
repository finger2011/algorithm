package week10onclass

// https://leetcode-cn.com/problems/counting-bits/
// leetcode 338. 比特位计数

func countBits(n int) []int {
	var bits = make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}
