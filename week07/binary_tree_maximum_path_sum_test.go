package week07

import (
	"finger2011/algggorithm/internel"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
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
				root: internel.CreateTreeByIntLevel([]int{1, 2, 3}),
			},
			want: 6,
		},
		{
			name: "test02",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{-10, 9, 20, null, null, 15, 7}),
			},
			want: 42,
		},
		{
			name: "test03",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{-6, null, 3, 2}),
			},
			want: 5,
		},
		{
			name: "test04",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{1, 2}),
			},
			want: 3,
		},
		{
			name: "test05",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{-3, -2, -1}),
			},
			want: -1,
		},
		{
			name: "test06",
			args: args{
				root: internel.CreateTreeByIntLevel([]int{-3}),
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
