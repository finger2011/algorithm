package week08onclass

// https://www.acwing.com/problem/content/852/
// acwing 850. Dijkstra求最短路 II

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
//     inputs := strings.Split(strings.Trim(input, " "), " ")
//     var m, n int
//     n, _ = strconv.Atoi(strings.Trim(inputs[0], "\n"))
//     m, _ = strconv.Atoi(strings.Trim(inputs[1], "\n"))
//     // fmt.Printf("m:%d, n:%d\n", m, n)
//     var times = make([][]int, m)
//     for k := 0; k < m; k++  {
//         input, _ = f.ReadString('\n')
//         var time []int
//         inputs := strings.Split(strings.Trim(input, " "), " ")
//         var num int
//         for i := 0; i < 3; i++ {
//             num, _ = strconv.Atoi(strings.Trim(inputs[i], "\n"))
//             time = append(time, num)
//         }
//         times[k] = time
//     }
//     // fmt.Printf("times:%v", times)
//     fmt.Printf("%d", networkDelayTime(times, n, 1))
// }


// type heap struct {
// 	heapList [][]int
// }
// func (m *heap) Push(val []int) error {
// 	m.heapList = append(m.heapList, val)
// 	return m.heapifyUp(len(m.heapList) - 1)
// }

// func (m *heap) heapifyUp(i int) error {
// 	for i > 1 {
// 		if m.heapList[i][0] < m.heapList[i/2][0] {
// 			var tmp = m.heapList[i]
// 			m.heapList[i] = m.heapList[i/2]
// 			m.heapList[i/2] = tmp
// 			i /= 2
// 		} else {
// 			break
// 		}
// 	}
// 	return nil
// }

// func (m *heap) Size() int {
// 	return len(m.heapList) - 1
// }

// func (m *heap) Pop() []int {
// 	var v = m.heapList[1]
// 	m.heapList[1] = m.heapList[len(m.heapList)-1]
// 	m.heapList = m.heapList[0 : len(m.heapList)-1]
// 	m.heapifyDown(1)
// 	return v
// }

// func (m *heap) heapifyDown(i int) {
// 	var ch = i * 2
// 	for ch < len(m.heapList) {
// 		if ch+1 < len(m.heapList) && m.heapList[ch+1][0] < m.heapList[ch][0] {
// 				ch++
// 		}
// 		if m.heapList[i][0] > m.heapList[ch][0] {
// 			var tmp = m.heapList[i]
// 			m.heapList[i] = m.heapList[ch]
// 			m.heapList[ch] = tmp
// 			i = ch
// 			ch = i * 2
// 		} else {
// 			break
// 		}
// 	}
// 	return 
// }

// //带堆的dijkstra
// func networkDelayTime(times [][]int, n int, k int) int {
// 	dist := make([]int, n + 1)
// 	for i := 0; i <= n; i++ {
// 		dist[i] = 1e9
// 	}
// 	dist[k] = 0
// 	//出边数组
// 	ver := make([][]int, n + 1)
// 	edge := make([][]int, n + 1)
// 	expand := make([]bool, n + 1)
// 	for i := 0; i < len(times); i++ {
// 		x, y, z := times[i][0], times[i][1], times[i][2]
// 		ver[x] = append(ver[x], y)
// 		edge[x] = append(edge[x], z)
// 	}
// 	h := heap{
// 		heapList:make([][]int, 1, n + 1),
// 	}
// 	h.Push([]int{0, k})
// 	for h.Size() > 0 {
// 		x := h.Pop()[1]
// 		if expand[x] {
// 			continue
// 		}
// 		expand[x] = true
// 		for i := 0; i < len(ver[x]); i++ {
// 			y, z := ver[x][i], edge[x][i]
// 			if dist[y] > dist[x] + z {
// 				dist[y] = dist[x] + z
// 				h.Push([]int{dist[y], y})
// 			}
// 		}
// 	}
// 	if dist[n] == 1e9 {
// 	    return -1
// 	}
// 	return dist[n]
// }
