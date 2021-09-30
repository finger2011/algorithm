package week01onclass

import (
	"strconv"
)

//https://leetcode-cn.com/problems/basic-calculator-ii/
//leetcode 227 基本计算器 II （无括号版本）

func calculateWithoutParenthess(s string) int {
	var opStack []string
	var param string
	var stack []string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "+" || string(s[i]) == "-" || string(s[i]) == "*" || string(s[i]) == "/" {
			stack = append(stack, param)
			param = ""
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
		} else if string(s[i]) != " " {
			param += string(s[i])
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
	// fmt.Printf("stack:%v", stack)
	return evalRPN(stack)
}

func evalRPN(tokens []string) int {
	var ant int
	var stack []string
	for i := 0; i < len(tokens); i++ {
		if string(tokens[i]) == "+" || string(tokens[i]) == "-" || string(tokens[i]) == "*" || string(tokens[i]) == "/" {
			var a, _ = strconv.Atoi(stack[len(stack)-1])
			var b, _ = strconv.Atoi(stack[len(stack)-2])
			stack = stack[0:(len(stack) - 2)]
			stack = append(stack, strconv.Itoa(cal(b, a, string(tokens[i]))))
		} else {
			stack = append(stack, string(tokens[i]))
		}
	}
	ant, _ = strconv.Atoi(stack[0])
	return ant
}

func getRank(op string) int {
	if op == "+" || op == "-" {
		return 1
	}
	if op == "*" || op == "/" {
		return 2
	}
	return 0
}

// func cal(a, b int, op string) int {
// 	if op == "+" {
// 		return a + b
// 	}
// 	if op == "-" {
// 		return a - b
// 	}
// 	if op == "*" {
// 		return a * b
// 	}
// 	if op == "/" {
// 		return a / b
// 	}
// 	return 0
// }
