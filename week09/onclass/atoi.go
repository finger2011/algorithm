package week09onclass

// https://leetcode-cn.com/problems/string-to-integer-atoi/
// leetcode 8. 字符串转换整数 (atoi)

func myAtoi(s string) int {
	start := 0
	// 1. 读入字符串并丢弃无用的前导空格
	for start < len(s) && s[start] == ' ' {
		start++
	}
	// 2. 检查正负号
	var negative = 1
	if start < len(s) && (s[start] == '-' || s[start] == '+') {
		if s[start] == '-' {
			negative = -1
		}
		start++
	}
	// 3. 读入数字
	var ans int
	var maxInt = 1<<31 - 1
	for start < len(s) && s[start] >= 48 && s[start] <= 57 {
		// ans*10 + int(s[start] - '0') > 1<<31 - 1
		if ans > (maxInt-int(s[start]-'0'))/10 {
			if negative == 1 {
				return maxInt
			}
			return -(1 << 31)
		}
		ans = ans*10 + int(s[start]-'0')
		start++

	}
	return ans * negative
}
