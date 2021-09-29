package week01onclass

import (
	"testing"
)

func createNodeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var guard = &ListNode{}

	var cur = guard
	for i := 0; i < len(nums); i++ {
		var node = &ListNode{Val: nums[i]}
		cur.Next = node
		cur = node
	}
	return guard.Next
}

//判断2个nodelist是否相同
func nodeListEqual(node1, node2 *ListNode) bool {
	for {
		if node1 == nil {
			return node2 == nil
		}
		if node2 == nil {
			return node1 == nil
		}
		if node1.Val != node2.Val {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
}

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				head: createNodeList([]int{1, 2, 3, 4, 5}),
			},
			want: createNodeList([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "test02",
			args: args{
				head: createNodeList([]int{1, 2}),
			},
			want: createNodeList([]int{2, 1}),
		},
		{
			name: "test03",
			args: args{
				head: createNodeList([]int{}),
			},
			want: createNodeList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !nodeListEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
