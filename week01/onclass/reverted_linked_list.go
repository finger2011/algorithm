package week01onclass

//https://leetcode-cn.com/problems/reverse-linked-list/
//leetcode 206

//ListNode list node
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var cur = head
	var prev *ListNode
	for {
		next := cur.Next
		cur.Next = prev
		if next == nil {
			break
		}
		prev = cur
		cur = next
	}
	return cur
}
