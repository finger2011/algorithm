package week09

// https://leetcode-cn.com/problems/reverse-only-letters/
// leetcode 917. 仅仅反转字母

func reverseOnlyLetters(s string) string {
	l, r := 0, len(s) - 1
	var ans = []byte(s)
	for l < r {
		for l < r && (!checkChar(ans[l])) {
			l++
		}
		for l < r && (!checkChar(ans[r])) {
			r--
		}
		tmp := ans[l]
		ans[l] = ans[r]
		ans[r] = tmp
		l++
		r--
	}
	return string(ans)
}

func checkChar(ch byte) bool {
	return (ch > 64 && ch < 91) || (ch > 96 && ch < 123)
}