package week03onclass

// https://leetcode-cn.com/problems/generate-parentheses/
// 22. 括号生成

var ansMap = make(map[int][]string)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if _, ok := ansMap[n]; ok {
		return ansMap[n]
	}
	var ans []string
	for i := 1; i <= n; i++ {
		var prev = generateParenthesis(i - 1)
		var next = generateParenthesis(n - i)
		for _, pre := range prev {
			for _, nex := range next {
				ans = append(ans, "("+pre+")"+nex)
			}
		}
	}
	ansMap[n] = ans
	return ans
}
