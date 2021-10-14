package week03onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// leetcode 236. 二叉树的最近公共祖先

var pNode, qNode, ansNode *internel.TreeNode

func lowestCommonAncestor(root, p, q *internel.TreeNode) *internel.TreeNode {
	pNode = p
	qNode = q
	ansNode = nil
	bfs(root)
	return ansNode
}

func bfs(root *internel.TreeNode) (bool, bool) {
	if root == nil {
		return false, false
	}
	left1, left2 := bfs(root.Left)
	right1, right2 := bfs(root.Right)
	var pContain = left1 || right1 || root.Val == pNode.Val
	var qContain = left2 || right2 || root.Val == qNode.Val
	if pContain && qContain && ansNode == nil {
		ansNode = root
	}
	return pContain, qContain
}
