package week10onclass

// https://www.acwing.com/problem/content/description/244/
// leetcode 243. 一个简单的整数问题2

// package main

// import (
//     "bufio"
//     "strconv"
//     "strings"
//     "os"
//     "fmt"
// )

// func main(){
//     f := bufio.NewReader(os.Stdin)
//     var input string
//     input, _ = f.ReadString('\n')
//     inputs := strings.Split(strings.Trim(input, "\n"), " ")
//     var m, n int
//     n, _ = strconv.Atoi(strings.Trim(inputs[0], "\n"))
//     m, _ = strconv.Atoi(strings.Trim(inputs[1], "\n"))
//     // fmt.Printf("m:%d, n:%d\n", m, n)
//     var nums = make([]int, n)
//     input, _ = f.ReadString('\n')
//     inputs = strings.Split(strings.Trim(input, "\n"), " ")
//     for i := 0; i < n; i++ {
//         nums[i], _ = strconv.Atoi(strings.Trim(inputs[i], "\n"))
//     }
//     // fmt.Printf("nums:%d\n", nums)
//     var tree = newSegmentTree(nums)
//     //读入m条指令
//     for m > 0 {
//         input, _ = f.ReadString('\n')
//         inputs = strings.Split(strings.Trim(input, "\n"), " ")
//         cal := inputs[0]
//         if cal == "Q" {
//             var n1, n2 int
//             n1, _ = strconv.Atoi(strings.Trim(inputs[1], "\n"))
//             n2, _ = strconv.Atoi(strings.Trim(inputs[2], "\n"))
//             // fmt.Printf("call:%s, %d, %d\n", cal, n1, n2)
//             fmt.Printf("%d\n", tree.getRank(1, n1, n2))
//         } else if cal == "C" {
//             var n1, n2, n3 int
//             n1, _ = strconv.Atoi(strings.Trim(inputs[1], "\n"))
//             n2, _ = strconv.Atoi(strings.Trim(inputs[2], "\n"))
//             n3, _ = strconv.Atoi(strings.Trim(inputs[3], "\n"))
//             // fmt.Printf("call:%s, %d, %d, %d\n", cal, n1, n2, n3)
//             tree.change(1, n1, n2, n3)
//         }
//         m--
//     }
// }

//线段树
type segmentTreeNode struct {
	l    int
	r    int
	val  int
	mask int
}

type segmentTree []segmentTreeNode

func newSegmentTree(nums []int) segmentTree {
	n := len(nums)
	tree := make(segmentTree, 4*n+1)
	tree.build(nums, 1, 1, n)
	return tree
}

func (t segmentTree) build(nums []int, p, l, r int) {
	t[p].l = l
	t[p].r = r
	if l == r {
		t[p].val = nums[l-1]
		return
	}
	mid := (l + r) / 2
	t.build(nums, p*2, l, mid)
	t.build(nums, p*2+1, mid+1, r)
	t[p].val = t[p*2].val + t[p*2+1].val
	return
}

func (t segmentTree) spead(p int) {
	if t[p].mask != 0 {
		t[p*2].val += t[p].mask * (t[p*2].r - t[p*2].l + 1)
		t[p*2].mask += t[p].mask
		t[p*2+1].val += t[p].mask * (t[p*2+1].r - t[p*2+1].l + 1)
		t[p*2+1].mask += t[p].mask
		t[p].mask = 0
	}
}

func (t segmentTree) change(p, l, r, val int) {
	if l <= t[p].l && r >= t[p].r {
		t[p].val += val * (t[p].r - t[p].l + 1)
		t[p].mask += val
		return
	}
	t.spead(p)
	mid := (t[p].l + t[p].r) / 2
	if l <= mid {
		t.change(p*2, l, r, val)
	}
	if r > mid {
		t.change(p*2+1, l, r, val)
	}
	t[p].val = t[p*2].val + t[p*2+1].val
}

func (t segmentTree) getRank(p, l, r int) int {
	if l <= t[p].l && r >= t[p].r {
		return t[p].val
	}
	t.spead(p)
	var ans int
	mid := (t[p].l + t[p].r) / 2
	if l <= mid {
		ans += t.getRank(p*2, l, r)
	}
	if r > mid {
		ans += t.getRank(p*2+1, l, r)
	}

	return ans
}
