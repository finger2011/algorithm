package week09

// https://leetcode-cn.com/problems/isomorphic-strings/
// leetcode 205. 同构字符串

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cMap := make(map[byte]byte, len(s))
	tMap := make(map[byte]byte, len(s))
	for i := 0; i < len(s); i++ {
		if _, ok := tMap[t[i]]; ok && s[i] != tMap[t[i]] {
			return false
		}
		if _, ok := cMap[s[i]]; !ok {
			cMap[s[i]] = t[i]
			tMap[t[i]] = s[i]
		} else if cMap[s[i]] != t[i] {
			return false
		}
		
	}
	return true
}
