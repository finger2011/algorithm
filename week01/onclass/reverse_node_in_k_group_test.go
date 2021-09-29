package week01onclass

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
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
				head: createNodeList([]int{1,2,3,4,5}),
				k:    2,
			},
			want: createNodeList([]int{2,1,4,3,5}),
		},
		{
			name: "test02",
			args: args{
				head: createNodeList([]int{1,2,3,4,5}),
				k:    3,
			},
			want: createNodeList([]int{3,2,1,4,5}),
		},
		{
			name: "test01",
			args: args{
				head: createNodeList([]int{1,2,3,4,5}),
				k:    1,
			},
			want: createNodeList([]int{1,2,3,4,5}),
		},
		{
			name: "test01",
			args: args{
				head: createNodeList([]int{1}),
				k:    1,
			},
			want: createNodeList([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !nodeListEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
