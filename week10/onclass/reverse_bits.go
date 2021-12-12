package week10onclass

// https://leetcode-cn.com/problems/reverse-bits/
// leetcode 190. 颠倒二进制位

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans = (ans << 1) | (num >> i & 1)
	}
	return ans
}
