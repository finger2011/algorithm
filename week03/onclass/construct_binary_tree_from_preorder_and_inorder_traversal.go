package week03onclass

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// leetcode 105. 从前序与中序遍历序列构造二叉树

func buildTree(preorder []int, inorder []int) *internel.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var root = &internel.TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	var leftPreorder, rightPreorder, leftInorder, rightInorder []int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			leftInorder = inorder[0:i]
			if i < len(inorder)-1 {
				rightInorder = inorder[i+1:]
			}
			break
		}
	}
	leftPreorder = preorder[1 : 1+len(leftInorder)]
	rightPreorder = preorder[1+len(leftInorder):]
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return root
}
