package week09

import "strconv"

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
// leetcode 438. 找到字符串中所有字母异位词

func findAnagrams(s string, p string) []int {
	var ans []int
	pHash := calAnagramsHash(p)
	for i := 0; i <= len(s)-len(p); i++ {
		t := calAnagramsHash(s[i : i+len(p)])
		if t == pHash {
			ans = append(ans, i)
		}
	}
	return ans
}

func calAnagramsHash(s string) string {
	var str = make([]int, 26)
	for i := 0; i < len(s); i++ {
		str[int(s[i]-97)]++
	}
	var ant string
	for i := 0; i < len(str); i++ {
		if str[i] > 0 {
			ant += strconv.Itoa(i) + strconv.Itoa(str[i]) + ","
		}
	}
	return ant
}
