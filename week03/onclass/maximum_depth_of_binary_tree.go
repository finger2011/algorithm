package week03onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
// leetcode 104. 二叉树的最大深度

func maxDepth(root *internel.TreeNode) int {
	if root == nil {
		return 0
	}
	return internel.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

//自底向上
var treeDeep int
var ans int
func maxDepth2(root *internel.TreeNode) int {
	if root == nil {
		return 0
	}
	ans = 0
	treeDeep = 1
	maxDeep(root)
	return ans
}

func maxDeep(root *internel.TreeNode)  {
	if root == nil {
		return 
	}
	ans = internel.Max(ans, treeDeep)
	treeDeep++
	maxDeep(root.Left)
	maxDeep(root.Right)
	treeDeep--
	return
}
