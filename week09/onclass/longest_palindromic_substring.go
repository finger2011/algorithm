package week09onclass

// https://leetcode-cn.com/problems/longest-palindromic-substring/
// leetcode 5. 最长回文子串

var preHash []int64
var nextHash []int64
var pow []int64
var b int64
var p int64

func longestPalindrome2(s string) string {
	var ansLen, ansStart int
	b, p = int64(131), int64(1e9 + 7)
	preHash = make([]int64, len(s) + 2)
	nextHash = make([]int64, len(s) + 1)
	pow =  make([]int64, len(s) + 1)
	pow[0] = 1
	for i := 1 ; i <= len(s); i++ {
		nextHash[i] = (nextHash[i - 1] * b + int64(s[i - 1] - 'a' + 1)) % p
		pow[i] = pow[i - 1] * b % p
	}
	for i := len(s); i >= 1; i-- {
		preHash[i] = (preHash[i + 1] * b + int64(s[i - 1] - 'a' + 1)) % p
	}
	for c := 0; c < len(s); c++ {
		left, right := 0, len(s) - 1 - c
		if c < right {
			right = c
		}
		for left < right {
			mid := (left + right + 1) / 2
			v1, v2 := calHash(c - mid, c + mid), calReverseHash(c - mid, c + mid)
			if v1 == v2 {
				left = mid
			} else {
				right = mid - 1
			}
		}
		if 2 * right + 1 > ansLen {
			ansLen = 2 * right + 1
			ansStart = c - right 
		}
	}
	for c := 1; c < len(s); c++ {
		left, right := -1, len(s) - 1 - c
		if c - 1 < right {
			right = c - 1
		}
		for left < right {
			mid := (left + right + 1) / 2
			if calHash(c - 1 - mid, c + mid) == calReverseHash(c - 1 - mid, c + mid) {
				left = mid
			} else {
				right = mid - 1
			}
		}
		if 2 * right + 2 > ansLen {
			ansLen = 2 * right + 2
			ansStart = c - 1 - right 
		}
	} 
	return s[ansStart:ansStart + ansLen]
}

func calHash(l, r int) int64 {
	return ((nextHash[r + 1] - nextHash[l] * pow[r - l + 1]) % p + p ) % p
}
func calReverseHash(l, r int) int64 {
	return ((preHash[l + 1] - preHash[r + 2] * pow[r - l + 1]) % p + p ) % p
}

func longestPalindrome(s string) string {
	var ansLen, ansStart int
	for c := 0; c < len(s); c++ {
		l, r := c - 1, c + 1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++ 
		}
		if r - l - 1 > ansLen {
			ansLen = r - l - 1
			ansStart = l + 1
		}
	}
	for c := 1; c < len(s); c++ {
		l, r := c - 1, c
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++ 
		}
		if r - l - 1 > ansLen {
			ansLen = r - l - 1
			ansStart = l + 1
		}
	}
	return s[ansStart:ansStart + ansLen]
}