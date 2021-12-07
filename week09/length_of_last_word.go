package week09

// https://leetcode-cn.com/problems/length-of-last-word/
// leetcode 58. 最后一个单词的长度

func lengthOfLastWord(s string) int {
	start := len(s) - 1
	for start >= 0 && s[start] == ' ' {
		start--
	}
	var ans int
	for start >= 0 && s[start] != ' ' {
		start--
		ans++
	}
	return ans
}