package week08onclass

// https://leetcode-cn.com/problems/network-delay-time/
// leetcode 743. 网络延迟时间

type heap struct {
	heapList [][]int
}
func (m *heap) Push(val []int) error {
	m.heapList = append(m.heapList, val)
	return m.heapifyUp(len(m.heapList) - 1)
}

func (m *heap) heapifyUp(i int) error {
	for i > 1 {
		if m.heapList[i][0] < m.heapList[i/2][0] {
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

func (m *heap) Pop() []int {
	var v = m.heapList[1]
	m.heapList[1] = m.heapList[len(m.heapList)-1]
	m.heapList = m.heapList[0 : len(m.heapList)-1]
	m.heapifyDown(1)
	return v
}

func (m *heap) heapifyDown(i int) {
	var ch = i * 2
	for ch < len(m.heapList) {
		if ch+1 < len(m.heapList) && m.heapList[ch+1][0] < m.heapList[ch][0] {
				ch++
		}
		if m.heapList[i][0] > m.heapList[ch][0] {
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

//带堆的dijkstra
func networkDelayTime3(times [][]int, n int, k int) int {
	dist := make([]int, n + 1)
	for i := 0; i <= n; i++ {
		dist[i] = 1e9
	}
	dist[k] = 0
	//出边数组
	ver := make([][]int, n + 1)
	edge := make([][]int, n + 1)
	expand := make([]bool, n + 1)
	for i := 0; i < len(times); i++ {
		x, y, z := times[i][0], times[i][1], times[i][2]
		ver[x] = append(ver[x], y)
		edge[x] = append(edge[x], z)
	}
	h := heap{
		heapList:make([][]int, 1, n + 1),
	}
	h.Push([]int{0, k})
	for h.Size() > 0 {
		x := h.Pop()[1]
		if expand[x] {
			continue
		}
		expand[x] = true
		for i := 0; i < len(ver[x]); i++ {
			y, z := ver[x][i], edge[x][i]
			if dist[y] > dist[x] + z {
				dist[y] = dist[x] + z
				h.Push([]int{dist[y], y})
			}
		}
	}
	var ans int
	for i := 1; i <= n; i++ {
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	if ans == 1e9 {
		return -1
	}
	return ans
}

// 不带堆的dijkstra
func networkDelayTime2(times [][]int, n int, k int) int {
	dist := make([]int, n + 1)
	for i := 0; i <= n; i++ {
		dist[i] = 1e9
	}
	dist[k] = 0
	//出边数组
	ver := make([][]int, n + 1)
	edge := make([][]int, n + 1)
	expand := make([]bool, n + 1)
	for i := 0; i < len(times); i++ {
		x, y, z := times[i][0], times[i][1], times[i][2]
		ver[x] = append(ver[x], y)
		edge[x] = append(edge[x], z)
	}
	for r := 1; r <= n - 1; r++ {
		var x, tmp int
		tmp = 1e9
		for i := 1; i <= n ; i++ {
			if !expand[i] && dist[i] < tmp {
				tmp = dist[i]
				x = i
			}
		}
		expand[x] = true
		for i := 0; i < len(ver[x]); i++ {
			y, z := ver[x][i], edge[x][i]
			if dist[y] > dist[x] + z {
				dist[y] = dist[x] + z
			}
		}
	}
	var ans int
	for i := 1; i <= n; i++ {
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	if ans == 1e9 {
		return -1
	}
	return ans
}

//bellman-ford
func networkDelayTime(times [][]int, n int, k int) int {
	dist := make([]int, n + 1)
	for i := 0; i <= n; i++ {
		dist[i] = 1e9
	}
	dist[k] = 0
	for r := 1; r <= n - 1; r++ {
		changed := false
		for i := 0; i < len(times); i++ {
			x, y, z := times[i][0], times[i][1], times[i][2]
			if dist[x] + z < dist[y] {
				dist[y] = dist[x] + z
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	var ans int
	for i := 1; i <= n; i++ {
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	if ans == 1e9 {
		return -1
	}
	return ans
}