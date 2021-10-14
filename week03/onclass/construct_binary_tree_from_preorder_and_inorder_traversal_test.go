package week03onclass

import (
	"finger2011/algggorithm/internel"
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
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
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: internel.CreateTreeByIntLevel([]int{-1}),
		},
		{
			name: "test01",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: internel.CreateTreeByIntLevel([]int{3, internel.MinInt, 9, 20, internel.MinInt, 15, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got.TreeToInt(), tt.want.TreeToInt()) {
				t.Errorf("buildTree() = %v, want %v", got.TreeToInt(), tt.want.TreeToInt())
			}
		})
	}
}
