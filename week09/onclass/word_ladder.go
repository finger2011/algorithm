package week09onclass

// https://leetcode-cn.com/problems/word-ladder/
// leetcode 127. 单词接龙

var wordMap map[string]bool

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap = make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordMap[word] = true
	}
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	wordDepBegin := make(map[string]int, len(wordList)+1)
	wordDepEnd := make(map[string]int, len(wordList)+1)
	wordDepBegin[beginWord] = 1
	wordDepEnd[endWord] = 1
	qBegin := []string{beginWord}
	qEnd := []string{endWord}
	var res int
	for len(qEnd) > 0 || len(qBegin) > 0 {
		res, qBegin = expand(qBegin, wordDepBegin, wordDepEnd)
		if res != -1 {
			return res
		}
		res, qEnd = expand(qEnd, wordDepEnd, wordDepBegin)
		if res != -1 {
			return res
		}
	}
	return 0
}

func expand(q []string, dep, otherDep map[string]int) (int, []string) {
	if len(q) == 0 {
		return -1, q
	}
	word := q[0]
	q = q[1:]
	for i := 0; i < len(word); i++ {
		for ch := byte('a'); ch <= byte('z'); ch++ {
			if ch == word[i] {
				continue
			}
			ns := word[0:i] + string(ch)
			if i < len(word)-1 {
				ns += word[i+1:]
			}
			if _, ok := wordMap[ns]; !ok {
				continue
			}
			if _, ok := dep[ns]; ok && dep[ns] > 0 {
				continue
			}	
			if _, ok := otherDep[ns]; ok && otherDep[ns] > 0 {
				return dep[word] + otherDep[ns], q
			}
			dep[ns] = dep[word] + 1
			q = append(q, ns)
		}
	}
	return -1, q
}
