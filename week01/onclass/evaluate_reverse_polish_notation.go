package week01onclass

import (
	"strconv"
)

//https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
// leetcode 150 逆波兰式表达式求值

func evalRPNWithNoRank(tokens []string) int {
    var ant int
    var stack []string
    for i := 0; i < len(tokens); i++ {
        if string(tokens[i]) == "+" || string(tokens[i]) == "-" || string(tokens[i]) == "*" || string(tokens[i]) == "/" {
            var a, _ = strconv.Atoi(stack[len(stack) - 1])
            var b, _ = strconv.Atoi(stack[len(stack) - 2])
            stack = stack[0:(len(stack) - 2)]
            stack = append(stack, strconv.Itoa(cal(b, a, string(tokens[i]))))
        } else {
            stack = append(stack, string(tokens[i]))
        }
    }
    ant, _ = strconv.Atoi(stack[0])
    return ant
}

func cal(a, b int, op string) int {
    if op == "+" {
        return a + b
    }
    if op == "-" {
        return a - b
    }
    if op == "*" {
        return a * b
    }
    if op == "/" {
        return a / b
    }
    return 0
}