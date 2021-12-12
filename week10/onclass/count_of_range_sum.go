package week10onclass

// https://leetcode-cn.com/problems/count-of-range-sum/
// leetcode 327. 区间和的个数
// 快排5w个数据直接timeout，没找出原因

var m int
var values []int64

func countRangeSum(nums []int, lower int, upper int) int{
	n := len(nums)
	sum := make([]int64, n + 1)
	for i := 1; i <= n ; i++ {
		sum[i] = sum[i - 1] + int64(nums[i - 1]) 
	} 
	values = make([]int64, 3 * (n + 1))
	for i := 0; i <= n; i++ {
		values[i * 3] = sum[i]
		values[i * 3 + 1] = sum[i] - int64(upper)
		values[i * 3 + 2] = sum[i] - int64(lower)
	}
	quickSort3(values, 0, len(values) - 1)
	m = 0
	for i := 0; i < len(values); i++ {
		if i == 0 || values[i] != values[i - 1] {
			values[m] = values[i]
			m++
		}
	}
	tree := newSegmentTree2(m)
	tree.add(1, getHashValue(sum[0]), 1)
	var ans int
	for i := 1; i <= n ; i++ {
		ans += tree.getRank(1, getHashValue(sum[i] - int64(upper)), getHashValue(sum[i] - int64(lower))) 
		tree.add(1, getHashValue(sum[i]), 1)
	}
	return ans
}



func getHashValue(val int64) int {
	left, right := 0, m - 1
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


func quickSort3(values []int64, left, right int) {
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
		quickSort3(values, left, p-1)
	}
	if right-p > 1 {
		quickSort3(values, p+1, right)
	}
}

//线段树
type segmentTreeNode2 struct {
	l int
	r int
	val int
}

type segmentTree2 []segmentTreeNode2

func newSegmentTree2(m int) segmentTree2 {
	tree := make(segmentTree2, 4 * m)
	tree.build(1, 0, m - 1)
	return tree
}

func (t segmentTree2) build( p, l, r int)  {
	t[p].l = l
	t[p].r = r
	if l == r {
		t[p].val = 0
		return
	}
	mid := (l + r) / 2
	t.build(p * 2, l, mid)
	t.build(p * 2 + 1, mid + 1, r)
	t[p].val = t[p * 2].val + t[p * 2 + 1].val
	return
}

func (t segmentTree2) add(p, x, val int)  {
	if t[p].l == t[p].r {
		t[p].val += val
		return
	}
	mid := (t[p].l + t[p].r) / 2
	if x <= mid {
		t.add(p * 2, x, val)
	} else {
		t.add(p * 2 + 1, x, val)
	}
	t[p].val = t[p * 2].val + t[p * 2 + 1].val
}

func (t segmentTree2) getRank(p, l, r int) int {
	if l <= t[p].l && r >= t[p].r {
		return t[p].val
	}
	var ans int
	mid := (t[p].l + t[p].r) / 2
	if l <= mid {
		ans += t.getRank(p * 2, l, r)
	} 
	if r > mid {
		ans += t.getRank(p * 2 + 1, l, r)
	}
	
	return ans
}