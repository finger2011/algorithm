package week09onclass

// https://leetcode-cn.com/problems/implement-strstr/
// leetcode 28. 实现 strStr()

func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	n, m := len(haystack), len(needle)
	haystack = " " + haystack
	needle = " " + needle
	next := make([]int, m + 1)
	for i, j := 2, 0; i <= m; i++ {
		for j > 0 && needle[j + 1] != needle[i] {
			j = next[j]
		}
		if needle[j + 1] == needle[i] {
			j++
		}
		next[i] = j
	}
	for i, j := 1, 0; i <= n; i++ {
		for j > 0 && (j == m || haystack[i] != needle[j + 1]) {
			j = next[j]
		}
		if haystack[i] == needle[j + 1] {
			j++
		}
		if j == m {
			return i - m
		}
	}
	return -1
}

func strStr(haystack string, needle string) int {
	b, p := int64(131), int64(1e9+7)
	hHash := make([]int64, len(haystack) + 1)
	for i := 1; i <= len(haystack); i++ {
		hHash[i] = (hHash[i - 1] * b + int64(haystack[i - 1] - 'a' + 1)) % p
	}
	nHash, pow := int64(0), int64(1)
	for i := 0; i < len(needle); i++ {
		pow = pow * b % p
		nHash = (nHash * b + int64(needle[i] - 'a' + 1)) % p
	}
	for l := 1; l <= len(haystack) - len(needle) + 1; l++ {
		r := l + len(needle) - 1
		if nHash == ((hHash[r] - hHash[l - 1] * pow) % p + p) % p {
			return l - 1
		}
	}
	return -1
}