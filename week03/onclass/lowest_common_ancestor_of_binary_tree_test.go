package week03onclass

import (
	"finger2011/algggorithm/internel"
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *internel.TreeNode
		p    *internel.TreeNode
		q    *internel.TreeNode
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
				root: internel.CreateTreeByIntLevel([]int{1, internel.MinInt, 2}),
				p:    &internel.TreeNode{Val: 1},
				q:    &internel.TreeNode{Val: 2},
			},
			want: &internel.TreeNode{Val: 1},
		},
		{
			name: "test02",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{3, 5, 1, 6, 2, 0, 8, internel.MinInt, 7, 4}),
				p:    &internel.TreeNode{Val: 5},
				q:    &internel.TreeNode{Val: 4},
			},
			want: &internel.TreeNode{Val: 5},
		},
		{
			name: "test03",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{3, 5, 1, 6, 2, 0, 8, internel.MinInt, 7, 4}),
				p:    &internel.TreeNode{Val: 5},
				q:    &internel.TreeNode{Val: 1},
			},
			want: &internel.TreeNode{Val: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
