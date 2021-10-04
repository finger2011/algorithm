package week02onclass

import (
	"fmt"
	"strconv"
)

//https://leetcode-cn.com/problems/group-anagrams/
// leetcode 49 字母异位词分组

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}
	var ant [][]string
	var strMap = make(map[string][]string, len(strs))
	for _, str := range strs {
		strMap[calGroupAnagramsHash(str)] = append(strMap[calGroupAnagramsHash(str)], str)
	}
	for _, str := range strMap {
		if len(str) > 0 {
			ant = append(ant, str)
		}
	}
	return ant
}

func calGroupAnagramsHash(s string) string {
	var str = make([]int, 26)
	for i := 0; i < len(s); i++ {
		str[int(s[i]-97)]++
	}
	var ant string
	for i := 0; i < len(str); i++ {
		if str[i] > 0 {
			ant += strconv.Itoa(i) + strconv.Itoa(str[i]) + ","
		}
	}
	fmt.Println("cal:", ant)
	return ant
}
