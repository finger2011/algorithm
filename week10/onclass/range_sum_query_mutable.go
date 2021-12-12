package week10onclass

// https://leetcode-cn.com/problems/range-sum-query-mutable/
// leetcode 307. 区域和检索 - 数组可修改

//NumArray 树状数组
type NumArray struct {
	a []int
	c []int
	n int
}

//Constructor 新建
func Constructor(nums []int) NumArray {
	n := len(nums)
	a := make([]int, n + 1)
	c := make([]int, n + 1)
	obj := NumArray{
		a:a,
		c:c,
		n:n,
	}
	for i := 1 ; i <= n; i++ {
		obj.a[i] = nums[i - 1]
		obj.add(i, a[i])
	}
	
	return obj
}

//Update 更新
func (n *NumArray) Update(index int, val int)  {
	index++
	n.add(index, val - n.a[index])
	n.a[index] = val
}

//SumRange 区间和
func (n *NumArray) SumRange(left int, right int) int {
	left++
	right++
	return n.query(right) - n.query(left - 1)
}

func (n *NumArray) query(x int) int {
	var ans int
	for ; x > 0; x -= lowbit(x) {
		ans += n.c[x]
	}
	return ans
}

func (n *NumArray) add(x, delta int)  {
	for ; x <= n.n; x += lowbit(x) {
		n.c[x] += delta
	}
}

func lowbit(x int) int {
	return x & -x
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */