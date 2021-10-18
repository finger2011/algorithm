package week04onclass

import (
	"finger2011/algggorithm/internel"
	"reflect"
	"testing"
)

func Test_inorderSuccessor(t *testing.T) {
	type args struct {
		root *internel.TreeNode
		p    *internel.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *internel.TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{2, 1, 3}),
				p:    &internel.TreeNode{Val: 1},
			},
			want: &internel.TreeNode{Val: 2},
		},
		{
			name: "test02",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{5,3,6,2,4,internel.MinInt,internel.MinInt,1}),
				p:    &internel.TreeNode{Val: 6},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inorderSuccessor(tt.args.root, tt.args.p)
			if (got == nil && tt.want != nil) || (got != nil && tt.want == nil) {
				t.Errorf("inorderSuccessor() = %v, want %v", got, tt.want)
			}
			if got != nil && tt.want != nil && !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("inorderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
