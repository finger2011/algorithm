package week09

// https://leetcode-cn.com/problems/reverse-string-ii/
// leetcode 541. 反转字符串 II

func reverseStr(s string, k int) string {
	var i, j int
	var t string
	for i = 2 * k; i < len(s); i += 2 * k {
		t += reverseString2([]byte(s[j : j+k])) + s[j+k:i]
		j = i
	}
	if len(s) - j >= k {
		t += reverseString2([]byte(s[j : j+k])) + s[j+k:]
	} else {
		t += reverseString2([]byte(s[j:]))
	}
	return t
}

func reverseString2(s []byte) string {
	l, r := 0, len(s)-1
	for l < r {
		var tmp = s[l]
		s[l] = s[r]
		s[r] = tmp
		l++
		r--
	}
	return string(s)
}
