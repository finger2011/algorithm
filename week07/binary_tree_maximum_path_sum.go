package week07

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
// leetcode 124. 二叉树中的最大路径和

// 存储每个节点最大路径（左右子树最多包含一个）
var maxPathSumTree map[*internel.TreeNode][]int

func maxPathSum(root *internel.TreeNode) int {
	maxPathSumTree = make(map[*internel.TreeNode][]int, 10)
	maxPathSumDfs(root)
	return maxPathSumTree[root][1]
}

func maxPathSumDfs(root *internel.TreeNode) {
	maxPathSumTree[root] = []int{-1e9, -1e9}
	if root == nil {
		return
	}
	maxPathSumDfs(root.Left)
	maxPathSumDfs(root.Right)
	maxPathSumTree[root][0] = internel.Max(maxPathSumTree[root.Left][0], maxPathSumTree[root.Right][0], 0) + root.Val
	maxPathSumTree[root][1] = internel.Max(
		maxPathSumTree[root.Left][1], 
		maxPathSumTree[root.Right][1], 
		maxPathSumTree[root.Left][0]+maxPathSumTree[root.Right][0]+root.Val,
		maxPathSumTree[root][0],
	)
	return
}
