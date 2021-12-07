package week09onclass

// https://leetcode-cn.com/problems/valid-palindrome/
// leetcode 125. 验证回文串
var pString string

func isPalindrome(s string) bool {
	pString = s
	l, r := getNextCh(0), getPreCh(len(s) - 1)
	for l < r {
		if !isCharEqual(s[l], s[r]) {
			return false
		}
		l = getNextCh(l + 1)
		r = getPreCh(r - 1)
	}
	return true
}

func getNextCh(i int) int {
	for i < len(pString) && !isDigOrLetter(pString[i]) {
		i++
	}
	return i
}

func getPreCh(i int) int {
	for i >= 0 && !isDigOrLetter(pString[i]) {
		i--
	}
	return i
}

func isDigOrLetter(ch byte) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isCharEqual(a, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a += 32
	}
	if b >= 'A' && b <= 'Z' {
		b += 32
	}
	return a == b
}