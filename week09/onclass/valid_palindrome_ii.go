package week09onclass

// https://leetcode-cn.com/problems/valid-palindrome-ii/
// leetcode 680. 验证回文字符串 Ⅱ

func validPalindrome(s string) bool {
	return validPalindromeWithTimes(s, 0, len(s) - 1, 1)
}

func validPalindromeWithTimes(s string, l, r, times int) bool {
	for l < r {
		if s[l] != s[r] {
			return times > 0 && (validPalindromeWithTimes(s, l + 1, r, times - 1) || validPalindromeWithTimes(s, l, r - 1, times - 1))
			
		}
		l++
		r--
	}
	return true
}