package week01onclass

//https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
//leetcode 25 K个一组翻转链表

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k <= 1{
        return head
    }
    var guard = new(ListNode)
    guard.Next = head
    var prev = guard
    for {
        var end = getEnd(head, k)
        if end == nil {
            break
        }
        var next = end.Next
        reverseKNode(head, k)
        prev.Next = end
        head.Next = next
        prev = head
        head = next
        if head == nil {
            break
        }
    }
    return guard.Next
}

func getEnd(node *ListNode, k int) *ListNode {
    for {
        k--
        if k == 0 {
            return node
        }
        if node.Next == nil {
            break
        }
        node = node.Next
    }
    return nil
}

func reverseKNode(head *ListNode, k int) {
    var prev = head
    head = head.Next
    k--
    for {
        var next = head.Next
        head.Next = prev
        k--
        if k == 0 {
            break
        }
        prev = head
        head = next      
    }
    return 
}