package week10

// https://www.acwing.com/problem/content/138/
// 136. 邻值查找


// package main

// import (
//     "fmt"
//     "bufio"
//     "strings"
//     "strconv"
//     "os"
// )

// func main(){
//     f := bufio.NewReader(os.Stdin)
//     var input string
//     input, _ = f.ReadString('\n')
//     var k int
//     k, _ = strconv.Atoi(string(input[0:(len(input) - 1)]))
//     // fmt.Printf("get k:%d\n", k)
//     var nums []int
//     var count = k
//     for count > 0 {
//         var input, _ = f.ReadString('\n')
//         input = strings.Trim(input, " ")
//         var inputs = strings.Split(input, " ")
//         for i := 0; i < len(inputs); i++ {
//             var num int 
//             num, _ = strconv.Atoi(string(inputs[i]))
//             nums = append(nums, num)
//             count--
//             if count <= 0 {
//                 break
//             }
//         }
//     }
//     var ant = lookup(k, nums)
//     for i := 0; i < len(ant); i++ {
//         fmt.Printf("%d %d\n", ant[i][0], ant[i][1])
//     }
// }

func lookup(k int, nums []int) [][]int {
	t := newAvlTree()
	var ans [][]int
	var numMap = make(map[int]int, k)
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			pre := t.getPrevVal(nums[i])
			next := t.getNextVal(nums[i])
			if pre != nil && next != nil {
				if nums[i] - pre.val <= next.val - nums[i] {
					ans = append(ans, []int{nums[i] - pre.val, numMap[pre.val]})
				} else {
					ans = append(ans, []int{next.val - nums[i], numMap[next.val]})
				}
			} else if pre != nil {
				ans = append(ans, []int{nums[i] - pre.val, numMap[pre.val]})
			} else if next != nil {
				ans = append(ans, []int{next.val - nums[i], numMap[next.val]})
			}
		}
		t = t.insert(nums[i])
		numMap[nums[i]] = i	+ 1
	}
	return ans
}