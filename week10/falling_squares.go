package week10

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/falling-squares/
// leetcode 699. 掉落的方块

var values []int

func fallingSquares(positions [][]int) []int {
	values = make([]int, len(positions) * 2)
	for i := 0; i < len(positions); i++ {
		values[2 * i] = positions[i][0]
		values[2 * i + 1] = positions[i][1] + positions[i][0] - 1
	}
	quickSort(values, 0, len(values) - 1)
	tree := newSegmentTree(len(values)) 
	var ans = make([]int, len(positions))
	for i := 0; i < len(positions); i++ {
		// 获取高度:重叠 + 本身高度
		left, right := getHashValue(positions[i][0]), getHashValue(positions[i][1] + positions[i][0] - 1)
		height := positions[i][1] + tree.getRank(1, left, right)
		tree.change(1, left, right, height)
		ans[i] = tree.getTop()
	}
	return ans
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

func getHashValue(val int) int {
	left, right := 0, len(values) - 1
	for left <= right {
		mid := (left + right) /2
		if values[mid] == val {
			return mid
		} else if values[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//线段树
type segmentTreeNode struct {
	l    int
	r    int
	val  int
	mask int
}

type segmentTree []segmentTreeNode

func newSegmentTree(n int) segmentTree {
	tree := make(segmentTree, 4*n)
	tree.build(1, 1, n)
	return tree
}

func (t segmentTree) build(p, l, r int) {
	t[p].l = l
	t[p].r = r
	if l == r {
		t[p].val = 0
		return
	}
	mid := (l + r) / 2
	t.build(p*2, l, mid)
	t.build(p*2+1, mid+1, r)
	t[p].val = internel.Max(t[p*2].val, t[p*2+1].val)
	return
}

func (t segmentTree) spead(p int) {
	if t[p].mask != 0 {
		t[p*2].val = t[p].mask
		t[p*2].mask = t[p].mask
		t[p*2+1].val = t[p].mask
		t[p*2+1].mask = t[p].mask
		t[p].mask = 0
	}
}

func (t segmentTree) change(p, l, r, val int) {
	if l <= t[p].l && r >= t[p].r {
		t[p].val = val
		t[p].mask = val
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
	t[p].val = internel.Max(t[p*2].val, t[p*2+1].val)
}

func (t segmentTree) getRank(p, l, r int) int {
	if l <= t[p].l && r >= t[p].r {
		return t[p].val
	}
	t.spead(p)
	var ans int
	mid := (t[p].l + t[p].r) / 2
	if l <= mid {
		ans = internel.Max(ans, t.getRank(p*2, l, r))
	}
	if r > mid {
		ans = internel.Max(ans, t.getRank(p*2+1, l, r))
	}

	return ans
}

func (t segmentTree) getTop() int {
	return t[1].val
}
