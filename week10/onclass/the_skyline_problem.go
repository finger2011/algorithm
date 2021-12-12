package week10onclass

// https://leetcode-cn.com/problems/the-skyline-problem/
// leetcode 218. 天际线问题

type event struct {
	pos    int
	index  int
	t      int
	height int
}

func getSkyline(buildings [][]int) [][]int {
	events := make([]event, 0)
	for i := 0; i < len(buildings); i++ {
		left, right, height := buildings[i][0], buildings[i][1], buildings[i][2]
		events = append(events, event{
			pos:    left,
			height: height,
			index:  i,
			t:      1,
		})
		events = append(events, event{
			pos:    right,
			height: height,
			index:  i,
			t:      -1,
		})
	}
	quickSort(events, 0, len(events) - 1)
	var ans [][]int
	removed := make([]bool, len(buildings))
	var q = heap{
		heapList:make([][]int, 1, 1),
	}
	for i := 0; i < len(events); i++ {
		if events[i].t == 1 {
			q.push([]int{events[i].height, events[i].index})
		} else {
			removed[events[i].index] = true
		}
		if i == len(events) - 1 || events[i].pos != events[i + 1].pos {
			for q.size() > 0 && removed[q.top()[1]] {
				q.pop()
			}
			height := 0 
			if q.size() > 0 {
				height = q.top()[0]
			}
			if len(ans) == 0 || height != ans[len(ans) - 1][1] {
				ans = append(ans, []int{events[i].pos, height})
			}
			
		}
	}
	return ans
}


type heap struct {
	heapList [][]int
}
func (m *heap) push(val []int) error {
	m.heapList = append(m.heapList, val)
	return m.heapifyUp(len(m.heapList) - 1)
}

func (m *heap) heapifyUp(i int) error {
	for i > 1 {
		if m.heapList[i][0] > m.heapList[i/2][0] {
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

func (m *heap) size() int {
	return len(m.heapList) - 1
}

func (m *heap) top() []int {
	return m.heapList[1]
}

func (m *heap) pop() []int {
	var v = m.heapList[1]
	m.heapList[1] = m.heapList[len(m.heapList)-1]
	m.heapList = m.heapList[0 : len(m.heapList)-1]
	m.heapifyDown(1)
	return v
}

func (m *heap) heapifyDown(i int) {
	var ch = i * 2
	for ch < len(m.heapList) {
		if ch+1 < len(m.heapList) && m.heapList[ch+1][0] > m.heapList[ch][0] {
				ch++
		}
		if m.heapList[i][0] < m.heapList[ch][0] {
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


func quickSort(values []event, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j].pos >= temp.pos {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i].pos <= temp.pos {
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
