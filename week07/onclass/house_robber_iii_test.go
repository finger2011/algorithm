package week07onclass

import (
	"finger2011/algggorithm/internel"
	"testing"
)

func Test_rob(t *testing.T) {
	var null = internel.MinInt
	type args struct {
		root *internel.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{3, 2, 3, null, 3, null, 1}),
			},
			want: 7,
		},
		{
			name: "test02",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{3, 4, 5, 1, 3, null, 1}),
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.root); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
