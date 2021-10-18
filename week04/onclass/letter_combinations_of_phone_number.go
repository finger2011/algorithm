package week04onclass

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
// 17. 电话号码的字母组合

var ans []string
var dig string

//DFS
func letterCombinationsByDFS(digits string) []string {
	ans = []string{}
	if len(digits) == 0 {
		return ans
	}
	dig = digits
	dfs(0, "")
	return ans
}

func dfs(index int, str string) {
	if index >= len(dig) {
		ans = append(ans, str)
		return
	}
	var codes = getContent(string(dig[index]))
	for i := 0; i < len(codes); i++ {
		dfs(index+1, str+codes[i])
	}
}

// 按序遍历
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var aut []string
	for i := 0; i < len(digits); i++ {
		var codes = getContent(string(digits[i]))
		if len(codes) == 0 {
			continue
		}
		if len(aut) == 0 {
			aut = codes
			continue
		}
		var j int
		var length = len(aut)
		for j = 0; j < length; j++ {
			for k := 0; k < len(codes); k++ {
				aut = append(aut, aut[j]+codes[k])
			}
		}
		aut = aut[j:]
	}
	return aut
}

func getContent(num string) []string {
	if num == "2" {
		return []string{"a", "b", "c"}
	}
	if num == "3" {
		return []string{"d", "e", "f"}
	}
	if num == "4" {
		return []string{"g", "h", "i"}
	}
	if num == "5" {
		return []string{"j", "k", "l"}
	}
	if num == "6" {
		return []string{"m", "n", "o"}
	}
	if num == "7" {
		return []string{"p", "q", "r", "s"}
	}
	if num == "8" {
		return []string{"t", "u", "v"}
	}
	if num == "9" {
		return []string{"w", "x", "y", "z"}
	}
	return []string{}
}
