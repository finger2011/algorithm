package week01onclass

//https://leetcode-cn.com/problems/linked-list-cycle-ii/
//leetcode 142 获取环形链表入环的第一个节点

func detectCycle(head *ListNode) *ListNode {
    var first = head
    var meetNode, err = hasCycle2(head)
    if err != nil {
		return nil
	}
	for {
		if first == meetNode {
			return meetNode
		}
		meetNode = meetNode.Next
		first = first.Next
	}
}