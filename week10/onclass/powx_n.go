package week10onclass

// https://leetcode-cn.com/problems/powx-n/
// leetcode 50. Pow(x, n)

func myPow(x float64, n int) float64 {
    if n == 0 {
		return 1.0
	}
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	var ans, tmp float64
	ans, tmp = 1, x
	for n > 0 {
		if n & 1 == 1 {
			ans *= tmp 
		}
		tmp *= tmp
		n = n >> 1 
	}
	return ans
}