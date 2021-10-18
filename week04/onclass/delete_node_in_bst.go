package week04onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/delete-node-in-a-bst/
// leetcode 450. 删除二叉搜索树中的节点

func deleteNode(root *internel.TreeNode, key int) *internel.TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} 
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		var next = root.Right
		for next.Left != nil {
			next = next.Left
		}
		root.Right = deleteNode(root.Right, next.Val)
		root.Val = next.Val
		return root
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
