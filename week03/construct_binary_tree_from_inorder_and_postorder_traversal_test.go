package week03

import (
	"finger2011/algggorithm/internel"
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
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
				inorder:   []int{9, 3, 15, 20, 7},
				postorder: []int{9, 15, 7, 20, 3},
			},
			want: internel.CreateTreeByIntLevel([]int{3, 9, 20, internel.MinInt, internel.MinInt, 15, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got.TreeToInt(), tt.want.TreeToInt()) {
				t.Errorf("buildTree() = %v, want %v", got.TreeToInt(), tt.want.TreeToInt())
			}
		})
	}
}
