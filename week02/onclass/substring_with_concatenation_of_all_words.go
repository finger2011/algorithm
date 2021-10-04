package week02onclass

// https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/
// leetcode 30 串联所有单词的子串

func findSubstring(s string, words []string) []int {
	var ant = []int{}
	var length int
	var wordMap = make(map[string]int, len(words))
	for i := 0; i < len(words); i++ {
		length += len(words[i])
		wordMap[words[i]]++
	}
	for i := 0; i+length <= len(s); i++ {
		if valid(s[i:i+length], len(words[0]), wordMap) {
			ant = append(ant, i)
		}
	}
	return ant
}

// 位移k
func findSubstring2(s string, words []string) []int {
	var ant = []int{}
	var length int
	var wordMap = make(map[string]int, len(words))

	for i := 0; i < len(words); i++ {
		length += len(words[i])
		wordMap[words[i]]++
	}
	for i := 0; i < len(words[0]); i++ {
		var splitMap = make(map[string]int, len(words))
		var k int
		for j := i; j+length <= len(s); j += len(words[0]) {
			if j == i {
				for k = j; k < j+length; k += len(words[0]) {
					splitMap[s[k:k+len(words[0])]]++
				}
			} else {
				splitMap[s[k:k+len(words[0])]]++
				k += len(words[0])
				splitMap[s[j-len(words[0]):j]]--
				if 0 == splitMap[s[j-len(words[0]):j]] {
					delete(splitMap, s[j-len(words[0]):j])
				}
			}
			if mapEqual(splitMap, wordMap) {
				ant = append(ant, j)
			}
		}

	}
	return ant
}

func valid(s string, length int, wordMap map[string]int) bool {
	var splitMap = make(map[string]int, len(s)/length)
	for i := 0; i < len(s); i += length {
		splitMap[s[i:i+length]]++
	}
	return mapEqual(wordMap, splitMap)
}

func mapEqual(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for i, val1 := range map1 {
		val2, ok := map2[i]
		if !ok || val1 != val2 {
			return false
		}
	}
	for i, val1 := range map2 {
		val2, ok := map1[i]
		if !ok || val1 != val2 {
			return false
		}
	}
	return true
}
