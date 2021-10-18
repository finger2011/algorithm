package week04

import (
	"finger2011/algggorithm/internel"
	"reflect"
	"testing"
)

func Test_convertBST(t *testing.T) {
	type args struct {
		root *internel.TreeNode
	}
	var null = internel.MinInt
	tests := []struct {
		name string
		args args
		want *internel.TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8}),
			},
			want: internel.CreateTreeByIntLevel([]int{30, 36, 21, 36, 35, 26, 15, null, null, null, 33, null, null, null, 8}),
		},
		{
			name: "test02",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{0, null, 1}),
			},
			want: internel.CreateTreeByIntLevel([]int{1, null, 1}),
		},
		{
			name: "test03",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{1, 0, 2}),
			},
			want: internel.CreateTreeByIntLevel([]int{3, 3, 2}),
		},
		{
			name: "test04",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{3, 2, 4, 1}),
			},
			want: internel.CreateTreeByIntLevel([]int{7, 9, 4, 10}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBST(tt.args.root); !reflect.DeepEqual(got.TreeToInt(), tt.want.TreeToInt()) {
				t.Errorf("convertBST() = %v, want %v", got.TreeToInt(), tt.want.TreeToInt())
			}
		})
	}
}
