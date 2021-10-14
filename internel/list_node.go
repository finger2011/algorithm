package internel

//ListNode ListNode define
type ListNode struct {
	Val  int
	Next *ListNode
}

//PrintListNode print []int
func (node *ListNode) PrintListNode() []int {
	var ans = []int{}
	for node != nil {
		ans = append(ans, node.Val)
		node = node.Next
	}
	return ans
}

//CreateListNode CreateListNode by []int
func CreateListNode(nums []int) *ListNode {
	return createListNode(nums)
}

func createListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var root = &ListNode{Val: nums[0]}
	var node = root
	for i := 1; i < len(nums); i++ {
		var newNode = &ListNode{Val: nums[i]}
		node.Next = newNode
		node = newNode
	}
	return root
}
