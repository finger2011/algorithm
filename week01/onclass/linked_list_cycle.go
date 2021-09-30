package week01onclass

import (
	"errors"
)

//https://leetcode-cn.com/problems/linked-list-cycle/
// leetcode 141 判断链表是否有环

var errNoCycle = errors.New("list has no cycle")

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    var fast = head
    for head != nil && fast != nil  {
        if fast.Next == nil {
            return false
        }
        head = head.Next
        fast = fast.Next.Next
        if head == fast {
            return true
        }
    }
	return false
}


func hasCycle2(head *ListNode) (*ListNode, error) {
    if head == nil {
        return head, nil
    }
    var fast = head
    for head != nil && fast != nil  {
        if fast.Next == nil {
            return nil, errNoCycle
        }
        head = head.Next
        fast = fast.Next.Next
        if head == fast {
            return fast, nil
        }
    }
	return nil, errNoCycle
}