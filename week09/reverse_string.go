package week09

// https://leetcode-cn.com/problems/reverse-string/
// leetcode 344. 反转字符串

func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		var tmp = s[l]
		s[l] = s[r]
		s[r] = tmp
		l++
		r--
	}
}
