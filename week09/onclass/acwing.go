package week09onclass

// https://www.acwing.com/problem/content/847/
// 845. 八数码 结果不对，预计是估值函数的问题？

// package main

// import (
//     "bufio"
//     "strings"
//     "os"
//     "fmt"
// )

// func main(){
//     f := bufio.NewReader(os.Stdin)
//     var input string
//     input, _ = f.ReadString('\n')
//     inputs := strings.Split(strings.Trim(input, "\n"), " ")
//     var start = strings.Join(inputs, "")
//     fmt.Printf("%d", eight(start))
// }

var h heap
var eDep map[string]int

type eval struct {
	e int
	s string
}

func eight(start string) int {
	target := "12345678x"
	h = heap{
		heapList: make([]eval, 1, 10),
	}
	h.Push(eval{e: eEval(start), s: start})
	eDep = make(map[string]int, 1)
	eDep[start] = 0
	for h.Size() > 0 {
		s := h.Pop()
		pos := findX(s.s)
		if pos >= 3 {
			eExpand(s.s, pos, pos-3)
		}
		if pos < 6 {
			eExpand(s.s, pos, pos+3)
		}
		if pos%3 != 0 {
			eExpand(s.s, pos, pos-1)
		}
		if pos%3 != 2 {
			eExpand(s.s, pos, pos+1)
		}
		if _, ok := eDep[target]; ok {
			return eDep[target]
		}
	}
	return -1
}

func eEval(s string) int {
	now := make([]int, 9)
	for i := 0; i < len(s); i++ {
		if s[i] != 'x' {
			now[int(s[i]-'0')] = i
		}
	}
	target := []int{0, 0, 1, 2, 3, 4, 5, 6, 7}
	var ans int
	for d := 1; d <= 8; d++ {
		nowx, nowy := now[d]/3, now[d]%3
		tx, ty := target[d]/3, target[d]%3
		ans += abs(nowx-tx) + abs(nowy-ty)
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func eExpand(s string, p1, p2 int) {
	var b = []byte(s)
	tmp := b[p1]
	b[p1] = b[p2]
	b[p2] = tmp
	if _, ok := eDep[string(b)]; ok {
		return
	}
	eDep[string(b)] = eDep[s] + 1
	h.Push(eval{
		e: eEval(string(b)),
		s: string(b),
	})
}

func findX(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == 'x' {
			return i
		}
	}
	return -1
}

type heap struct {
	heapList []eval
}

func (m *heap) Push(val eval) error {
	m.heapList = append(m.heapList, val)
	return m.heapifyUp(len(m.heapList) - 1)
}

func (m *heap) heapifyUp(i int) error {
	for i > 1 {
		if m.heapList[i].e < m.heapList[i/2].e {
			var tmp = m.heapList[i]
			m.heapList[i] = m.heapList[i/2]
			m.heapList[i/2] = tmp
			i /= 2
		} else {
			break
		}
	}
	return nil
}

func (m *heap) Size() int {
	return len(m.heapList) - 1
}

func (m *heap) Pop() eval {
	var v = m.heapList[1]
	m.heapList[1] = m.heapList[len(m.heapList)-1]
	m.heapList = m.heapList[0 : len(m.heapList)-1]
	m.heapifyDown(1)
	return v
}

func (m *heap) heapifyDown(i int) {
	var ch = i * 2
	for ch < len(m.heapList) {
		if ch+1 < len(m.heapList) && m.heapList[ch+1].e < m.heapList[ch].e {
			ch++
		}
		if m.heapList[i].e > m.heapList[ch].e {
			var tmp = m.heapList[i]
			m.heapList[i] = m.heapList[ch]
			m.heapList[ch] = tmp
			i = ch
			ch = i * 2
		} else {
			break
		}
	}
	return
}
