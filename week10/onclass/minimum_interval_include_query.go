package week10onclass

// https://leetcode-cn.com/problems/minimum-interval-to-include-each-query/
// leetcode 1851. 包含每个查询的最小区间

type event1 struct {
	pos int
	len int
	t   int
}

func minInterval(intervals [][]int, queries []int) []int {
	events := make([]event1, 0)
	for i := 0; i < len(intervals); i++ {
		len := intervals[i][1] - intervals[i][0] + 1
		events = append(events, event1{
			pos: intervals[i][0],
			len: len,
			t:   1,
		})
		events = append(events, event1{
			pos: intervals[i][1],
			len: len,
			t:   -1,
		})
	}
	for i := 0; i < len(queries); i++ {
		events = append(events, event1{
			pos: queries[i],
			len: i,
			t:   0,
		})
	}
	quickSort2(events, 0, len(events) - 1)
	var ans = make([]int, len(queries))
	var q = minHeap{
		heapList:make([]int, 1, 1),
	}
	removed := q
	for i := 0; i < len(events); i++ {
		switch events[i].t {
		case 1:
			q.push(events[i].len)
		case -1:
			removed.push(events[i].len)
		case 0: 
		for removed.size() > 0 &&removed.top() == q.top() {
			removed.pop()
			q.pop()
		}
		index := -1
		if q.size() > 0 {
			index = q.top()
		}
		ans[events[i].len] = index
		}
	}
	return ans
}


type minHeap struct {
	heapList []int
}
func (m *minHeap) push(val int) {
	m.heapList = append(m.heapList, val)
	m.heapifyUp(len(m.heapList) - 1)
	return
}

func (m *minHeap) heapifyUp(i int) error {
	for i > 1 {
		if m.heapList[i] < m.heapList[i/2] {
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

func (m *minHeap) size() int {
	return len(m.heapList) - 1
}

func (m *minHeap) top() int {
	return m.heapList[1]
}

func (m *minHeap) pop() int {
	var v = m.heapList[1]
	m.heapList[1] = m.heapList[len(m.heapList)-1]
	m.heapList = m.heapList[0 : len(m.heapList)-1]
	m.heapifyDown(1)
	return v
}

func (m *minHeap) heapifyDown(i int) {
	var ch = i * 2
	for ch < len(m.heapList) {
		if ch+1 < len(m.heapList) && m.heapList[ch+1] < m.heapList[ch] {
				ch++
		}
		if m.heapList[i] > m.heapList[ch] {
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

func compare(e1, e2 event1) bool {
	return e1.pos < e2.pos || (e1.pos == e2.pos && e1.t > e2.t)
}

func quickSort2(values []event1, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && (!compare(values[j], temp)) {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && (!compare(temp, values[i])) {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort2(values, left, p-1)
	}
	if right-p > 1 {
		quickSort2(values, p+1, right)
	}
}
