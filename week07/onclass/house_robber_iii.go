package week07onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/house-robber-iii/
// leetcode 337. 打家劫舍 III

var robTree map[*internel.TreeNode][]int

func rob(root *internel.TreeNode) int {
	robTree = make(map[*internel.TreeNode][]int, 10)
	dfsRob(root)
	return internel.Max(robTree[root][0], robTree[root][1])
}

func dfsRob(root *internel.TreeNode) {
	robTree[root] = []int{0, 0}
	if root == nil {
		return
	}
	dfsRob(root.Left)
	dfsRob(root.Right)
	robTree[root][0] = internel.Max(robTree[root.Left][0], robTree[root.Left][1]) +
		internel.Max(robTree[root.Right][0], robTree[root.Right][1])
	robTree[root][1] = robTree[root.Left][0] + robTree[root.Right][0] + root.Val
}
