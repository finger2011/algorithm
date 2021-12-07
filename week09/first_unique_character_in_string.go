package week09

// https://leetcode-cn.com/problems/first-unique-character-in-a-string/
// leetcode 387. 字符串中的第一个唯一字符

func firstUniqChar(s string) int {
	var sMap = make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if sMap[s[i]] == 1 {
			return i
		}
	}
	return -1
}