package week01onclass

// https://leetcode-cn.com/problems/basic-calculator/
// leetcode 224 基本计算器（带括号）
// 注意正负数

//https://leetcode-cn.com/problems/basic-calculator-iii/
// leetcode 772 基本计算器3 也可以解决

func calculateWithParenthess(s string) int {
	var opStack []string
	var param string
	var stack []string
	var needZero = true
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "+" || string(s[i]) == "-" || string(s[i]) == "*" || string(s[i]) == "/" {
			if param != "" {
				stack = append(stack, param)
				param = ""
			}
			if (string(s[i]) == "+" || string(s[i]) == "-") && needZero {
				//补零处理正负号
				stack = append(stack, "0")
			}
			for len(opStack) > 0 && getRank(opStack[len(opStack)-1]) >= getRank(string(s[i])) {
				stack = append(stack, opStack[len(opStack)-1])
				opStack = opStack[0:(len(opStack) - 1)]
			}
			if len(opStack) > 0 {
				var j = len(opStack) - 1
				for {
					if getRank(opStack[j]) >= getRank(string(s[i])) {
						stack = append(stack, opStack[j])
					} else {
						break
					}
					j--
					if j < 0 {
						break
					}
				}
				if j < len(opStack)-1 {
					opStack = opStack[0 : j+1]
				}
			}
			opStack = append(opStack, string(s[i]))
			needZero = true
		} else if string(s[i]) == "(" {
			needZero = true
			opStack = append(opStack, string(s[i]))
		} else if string(s[i]) == ")" {
			if param != "" {
				stack = append(stack, param)
				param = ""
			}
			var j int
			for j = len(opStack) - 1; j >= 0; j-- {
				if opStack[j] == "(" {
					break
				}
				stack = append(stack, opStack[j])
			}
			opStack = opStack[0:j]
			needZero = false
		} else if string(s[i]) != " " {
			param += string(s[i])
			needZero = false
		}
	}
	if param != "" {
		stack = append(stack, param)
	}
	if len(opStack) > 0 {
		for i := len(opStack) - 1; i >= 0; i-- {
			stack = append(stack, opStack[i])
		}
	}
	// fmt.Printf("stack:%v\n", stack)
	return evalRPN(stack)
}
