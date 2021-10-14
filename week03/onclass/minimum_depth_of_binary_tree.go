package week03onclass

import (
	"math"
	"finger2011/algggorithm/internel"
)

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
// leetcode 111. 二叉树的最小深度

//自顶向下
func minDepth(root *internel.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var minDeep = math.MaxInt32
	if root.Left != nil {
		minDeep = internel.Min(minDepth(root.Left), minDeep)
	}
	if root.Right != nil {
		minDeep = internel.Min(minDepth(root.Right), minDeep)
	}
	return minDeep + 1
}

//自底向上
var treeDeep2 int
var ans2 int
func minDepth2(root *internel.TreeNode) int {
	if root == nil {
		return 0
	}
	//初始化
	ans2 = math.MaxInt32
	treeDeep2 = 0
	deep(root)
	return ans2
}

func deep(root *internel.TreeNode)  {
	if root == nil {	
		return 
	}
	treeDeep2++
	if root.Left == nil && root.Right == nil {
		//到叶子节点时，判断该叶子节点深度是否为较小
		ans2 = internel.Min(ans2, treeDeep2)
	} else {
		deep(root.Left)
		deep(root.Right)
	}
	treeDeep2--
	return
}
