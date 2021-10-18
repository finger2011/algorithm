package week04onclass

import (
	"finger2011/algggorithm/internel"
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*internel.ListNode
	}
	tests := []struct {
		name string
		args args
		want *internel.ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				lists: []*internel.ListNode{
					internel.CreateListNode([]int{1, 4, 5}),
					internel.CreateListNode([]int{1, 3, 4}),
					internel.CreateListNode([]int{2, 6}),
				},
			},
			want: internel.CreateListNode([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name: "test02",
			args: args{
				lists: []*internel.ListNode{},
			},
			want: internel.CreateListNode([]int{}),
		},
		{
			name: "test03",
			args: args{
				lists: []*internel.ListNode{nil},
			},
			want: internel.CreateListNode([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got.PrintListNode(), tt.want.PrintListNode()) {
				t.Errorf("mergeKLists() = %v, want %v", got.PrintListNode(), tt.want.PrintListNode())
			}
		})
	}
}
