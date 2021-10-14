package week03onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// leetcode 94. 二叉树的中序遍历

func inorderTraversal(root *internel.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ans []int
	if root.Left != nil {
		ans = append(ans, inorderTraversal(root.Left)...)
	}
	ans = append(ans, root.Val)
	if root.Right != nil {
		ans = append(ans, inorderTraversal(root.Right)...)
	}
	return ans
}
