package week01onclass

// https://leetcode-cn.com/problems/valid-parentheses/
// leetcode 20 有效括号

func isValid(s string) bool {
	var stack []string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[" {
			stack = append(stack, string(s[i]))
		} else {
			if len(stack) == 0 {
				return false
			}
			if string(s[i]) == ")" && stack[len(stack)-1] != "(" {
				return false
			}
			if string(s[i]) == "}" && stack[len(stack)-1] != "{" {
				return false
			}
			if string(s[i]) == "]" && stack[len(stack)-1] != "[" {
				return false
			}
			stack = stack[0:(len(stack) - 1)]
		}
	}
	return len(stack) == 0
}
