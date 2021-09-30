package week01onclass

import (
	"fmt"
)

//https://www.acwing.com/problem/content/description/138/
//邻值查找
//给定一个长度为 n 的序列 A，A 中的数各不相同。
// 对于 A 中的每一个数 Ai，求：
// min1≤j<i|Ai−Aj|
// 以及令上式取到最小值的 j（记为 Pi）。若最小值点不唯一，则选择使 Aj 较小的那个。
// 输入格式
// 第一行输入整数 n，代表序列长度。
// 第二行输入 n 个整数A1…An,代表序列的具体数值，数值之间用空格隔开。
// 输出格式
// 输出共 n−1 行，每行输出两个整数，数值之间用空格隔开。
// 分别表示当 i 取 2∼n 时，对应的 min1≤j<i|Ai−Aj| 和 Pi 的值。
// 数据范围
// n≤105,|Ai|≤109

//输入样例：
// 3
// 1 5 3
// 输出样例：
// 4 1
// 2 1

//DoubleListNode 双端链表
type DoubleListNode struct {
	Val  int
	Pos  int
	Next *DoubleListNode
	Prev *DoubleListNode
}

func lookup(k int, nums []int) [][]int {
	var numAddr = make(map[int]int, k)
	fmt.Printf("k:%d\n", k)
	for i := 0; i < k; i++ {
		numAddr[nums[i]] = i
	}

	var sortNums = append([]int{}, nums...)
	quickSort(sortNums, 0, k-1)
	var head = new(DoubleListNode)

	var address = make(map[int]*DoubleListNode, k)

	for i := 0; i < k; i++ {
		var node = &DoubleListNode{
			Val: sortNums[i],
			Pos: numAddr[sortNums[i]],
		}
		head.Next = node
		node.Prev = head
		head = node
		address[sortNums[i]] = node
	}

	head.Next = new(DoubleListNode)
	head.Next.Prev = head
	var ant [][]int
	for i := k - 1; i > 0; i-- {
		var node = address[nums[i]]
		if node.Prev.Prev != nil && (node.Next.Next == nil || nums[i]-node.Prev.Val <= node.Next.Val-nums[i]) {
			ant = append(ant, []int{nums[i] - node.Prev.Val, node.Prev.Pos + 1})
		} else {
			ant = append(ant, []int{node.Next.Val - nums[i], node.Next.Pos + 1})
		}
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}

	return ant
}

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

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
//     // fmt.Printf("get ant:%d\n", ant)
//     for i := len(ant) - 1; i >= 0; i-- {
//         fmt.Printf("%d %d\n", ant[i][0], ant[i][1])
//     }
// }
