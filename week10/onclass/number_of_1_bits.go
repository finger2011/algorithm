package week10onclass

// https://leetcode-cn.com/problems/number-of-1-bits/
// leetcode 191. 位1的个数

func hammingWeight3(num uint32) int {
    var ans int
	for num > 0 {
		ans++
		num = num & (num - 1)
	}
	return ans
}

func hammingWeight2(num uint32) int {
    var ans int
	for i := 0; i < 32; i++ {
		if ((num >> i) & 1) == 1 {
			ans++
		}
	}
	return ans
}

func hammingWeight(num uint32) int {
    var ans int
	for num > 0 {
		if num & 1 == 1 {
			ans++
		}
		num = num >> 1
	}
	return ans
}