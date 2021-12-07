package week09

// https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/
// leetcode 557. 反转字符串中的单词 III

func reverseWords3(s string) string {
	var ans string
	var i int
	var word string
	for i < len(s) {
		word, i = getNextWord3(s, i)
		ans += reverseString3([]byte(word))
		end := i
		for end < len(s) && s[end] == ' ' {
			end++
		}
		if end > i {
			ans += s[i:end]
			i = end
		}
	}
	return ans
}

func getNextWord3(s string, i int) (string, int) {
	end := i
	for end < len(s) && s[end] != ' ' {
		end++
	}
	if end >= len(s) {
		return s[i:], end
	}
	return s[i:end], end
}

func reverseString3(s []byte) string {
	l, r := 0, len(s)-1
	for l < r {
		var tmp = s[l]
		s[l] = s[r]
		s[r] = tmp
		l++
		r--
	}
	return string(s)
}
