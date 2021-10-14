package week03

import "finger2011/algggorithm/internel"

// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// leetcode 106. 从中序与后序遍历序列构造二叉树

func buildTree(inorder []int, postorder []int) *internel.TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	var root = &internel.TreeNode{Val: postorder[len(postorder)-1]}
	if len(inorder) == 1 {
		return root
	}
	var leftInorder, rightInorder, leftPostorder, rightPostorder []int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			leftInorder = inorder[0:i]
			if i < len(inorder)-1 {
				rightInorder = inorder[i+1:]
			}
			break
		}
	}
	leftPostorder = postorder[0:len(leftInorder)]
	rightPostorder = postorder[len(leftInorder) : len(leftInorder)+len(rightInorder)]
	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}
