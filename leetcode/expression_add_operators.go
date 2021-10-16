package leetcode

import (
	"fmt"
	"strconv"
)

var ansNum string
var ansTarget int
var ans []string
var length int64

func addOperators(num string, target int) []string {
	ansNum = num
	ansTarget = target
	length = int64(len(num))
	ans = []string{}
	dfs(0, 0, 0, "")
	return ans
}

func dfs(start, cur, prev int64, context string) {
	if start >= length {
		if int(cur) == ansTarget {
			ans = append(ans, context)
		}
		return
	}
	if context == "3-4*5+62*37*4" {
		fmt.Printf("%s cur:%v, prev:%v\n", context, cur, prev)
	}
	if context == "3-4*5+62*37" {
		fmt.Printf("%s cur:%v, prev:%v\n", context, cur, prev)
	}
	if context == "3-4*5+62" {
		fmt.Printf("%s cur:%v, prev:%v\n", context, cur, prev)
	}
	for i := start; i < length; i++ {
		if i > start && string(ansNum[start]) == "0" {
			break
		}
		var next32 int
		if i < length-1 {
			next32, _ = strconv.Atoi(string(ansNum[start : i+1]))
		} else {
			next32, _ = strconv.Atoi(string(ansNum[start:]))
		}
		if start == 0 {
			dfs(i+1, int64(next32), int64(next32), context+strconv.Itoa(next32))
		} else {
			dfs(i+1, cur+int64(next32), int64(next32), context+"+"+strconv.Itoa(next32))
			dfs(i+1, cur-int64(next32), -int64(next32), context+"-"+strconv.Itoa(next32))
			dfs(i+1, cur-prev+prev*int64(next32), prev*int64(next32), context+"*"+strconv.Itoa(next32))
		}
	}
}
