package week09onclass

// https://leetcode-cn.com/problems/generate-parentheses/
// leetcode 22. 括号生成

var pStr string
var pN int
var pAns []string

func generateParenthesis(n int) []string {
	pStr = ""
	pN = n
	pAns = []string{}
	pDfs(0, 0, 0)
	return pAns
}

func pDfs(i, l, r int)  {
	if l > pN || r > pN || !isPValid(pStr) {
		return
	}
	if i == 2 * pN {
		if isPValid(pStr) {
			pAns = append(pAns, pStr)
		}
		return
	}
	pStr += "("
	pDfs(i + 1, l + 1, r)
	pStr = pStr[0:len(pStr) - 1]
	pStr += ")"
	pDfs(i + 1, l, r + 1)
	pStr = pStr[0:len(pStr) - 1]
}

func isPValid(s string) bool {
	n := 0
	for i := 0; i < len(s); i++ {
		if '(' == s[i] {
			n++
		} else {
			if n == 0 {
				return false
			}
			n--
		}
	}
	return true
}