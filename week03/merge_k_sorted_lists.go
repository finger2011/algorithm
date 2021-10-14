package week03

import (
	"finger2011/algggorithm/internel"
)

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// leetcode 23. 合并K个升序链表

func mergeKLists(lists []*internel.ListNode) *internel.ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		var root = &internel.ListNode{}
		var node = root
		var first = lists[0]
		var second = lists[1]
		for first != nil && second != nil {
			if first.Val <= second.Val {
				node.Next = first		
				first = first.Next
			} else {
				node.Next = second
				second = second.Next
			}
			node = node.Next
		}
		if first != nil {
			node.Next = first
		}
		if second != nil {
			node.Next = second
		}
		return root.Next
	}
	var left = mergeKLists(lists[0:len(lists) / 2]) 
	var right = mergeKLists(lists[len(lists)/2:])
	return mergeKLists([]*internel.ListNode{left, right})
}