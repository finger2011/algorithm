package week03onclass

import (
	"fmt"
	"strconv"
)

// https://leetcode-cn.com/problems/powx-n/
// leetcode 50. Pow(x, n)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	var tmp = myPow(x, n/2)
	tmp *= tmp
	if n%2 == 1 {
		tmp *= x
	}
	tmp, _ = strconv.ParseFloat(fmt.Sprintf("%.8f", tmp), 64)
	return tmp
}
