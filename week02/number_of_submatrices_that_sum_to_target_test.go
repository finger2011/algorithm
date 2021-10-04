package week02

import "testing"

func Test_numSubmatrixSumTarget(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "test01",
		// 	args: args{
		// 		matrix: [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}},
		// 		target: 0,
		// 	},
		// 	want: 4,
		// },
		{
			name: "test02",
			args: args{
				matrix: [][]int{{1, -1}, {-1, 1}},
				target: 0,
			},
			want: 5,
		},
		{
			name: "test03",
			args: args{
				matrix: [][]int{{904}},
				target: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubmatrixSumTarget(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("numSubmatrixSumTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
