package week09onclass

// https://leetcode-cn.com/problems/repeated-string-match/
// leetcode 686. 重复叠加字符串匹配

func repeatedStringMatch(a string, b string) int {
	var ans = 1
	var aa = a
	for len(aa) < len(b) {
		aa += a
		ans++
	}
	var k = checkMatch(aa, b)
	if k >= 0 {
		return ans
	}
	aa += a
	ans++
	k = checkMatch(aa, b)
	if k >= 0 {
		return ans
	}
	return -1
}

func checkMatch(a, b string) int {
	bi, p := int64(131), int64(1e9+7)
	nHash, pow := int64(0), int64(1)
	for i := 0; i < len(b); i++ {
		pow = pow * bi % p
		nHash = (nHash * bi + int64(b[i] - 'a' + 1)) % p
	}
	hHash := make([]int64, len(a) + 1)
	for i := 1; i <= len(a); i++ {
		hHash[i] = (hHash[i - 1] * bi + int64(a[i - 1] - 'a' + 1)) % p
	}
	for l := 1; l <= len(a) - len(b) + 1; l++ {
		r := l + len(b) - 1
		if nHash == ((hHash[r] - hHash[l - 1] * pow) % p + p) % p && b == a[l - 1:r] {
			return l - 1
		}
	}
	return -1
}