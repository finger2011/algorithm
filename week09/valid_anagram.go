package week09

// https://leetcode-cn.com/problems/valid-anagram/
// leetcode 242. 有效的字母异位词


func isAnagram(s string, t string) bool {
	sMap := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		sMap[t[i]]--
	}
	for _, count := range sMap {
		if count != 0 {
			return false
		}
	}
	return true
}