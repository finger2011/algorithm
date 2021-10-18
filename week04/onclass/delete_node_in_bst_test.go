package week04onclass

import (
	"finger2011/algggorithm/internel"
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *internel.TreeNode
		key  int
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
				root: internel.CreateTreeByIntLevel([]int{5, 3, 6, 2, 4, internel.MinInt, 7}),
				key:  3,
			},
			want: internel.CreateTreeByIntLevel([]int{5, 4, 6, 2, internel.MinInt, internel.MinInt, 7}),
		},
		{
			name: "test02",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{5, 3, 6, 2, 4, internel.MinInt, 7}),
				key:  0,
			},
			want: internel.CreateTreeByIntLevel([]int{5, 3, 6, 2, 4, internel.MinInt, 7}),
		},
		{
			name: "test03",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{}),
				key:  0,
			},
			want: internel.CreateTreeByIntLevel([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got.TreeToInt(), tt.want.TreeToInt()) {
				t.Errorf("deleteNode() = %v, want %v", got.TreeToInt(), tt.want.TreeToInt())
			}
		})
	}
}
