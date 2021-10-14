package week03onclass

import (
	"finger2011/algggorithm/internel"
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	type args struct {
		root *internel.NodeTree
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{root: internel.CreateNodeTree([]int{1, internel.MinInt, 3, 2, 4, internel.MinInt, 5, 6})},
			want: []int{1, 3, 5, 6, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
