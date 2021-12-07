package week09

// https://leetcode-cn.com/problems/reverse-words-in-a-string/
// leetcode 151. 翻转字符串里的单词

func reverseWords(s string) string {
	var strs []string
	var start int
	var word string
	for start < len(s) {
		word, start = getNextWord(s, start)
		if word != "" {
			strs = append(strs, word)
		}
	}
	var ans string
	for i := len(strs) - 1; i >= 0; i-- {
		ans += strs[i] + " "
	}
	return ans[0:len(ans) - 1]
}

func getNextWord(s string, start int) (string, int) {
	for start < len(s) && s[start] == ' ' {
		start++
	}
	end := start
	for end < len(s) && s[end] != ' '{
		end++
	}
	if end >= len(s) {
		return string(s[start:]), end + 1
	}
	return string(s[start:end]), end + 1
}