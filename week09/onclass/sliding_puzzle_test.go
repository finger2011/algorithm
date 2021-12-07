package week09onclass

import "testing"

func Test_slidingPuzzle(t *testing.T) {
	type args struct {
		board [][]int
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
				board: [][]int{
					{1, 2, 3},
					{4, 0, 5},
				},
			},
			want: 1,
		},
		{
			name: "test02",
			args: args{
				board: [][]int{
					{1, 2, 3},
					{5, 4, 0},
				},
			},
			want: -1,
		},
		{
			name: "test03",
			args: args{
				board: [][]int{
					{4, 1, 2},
					{5, 0, 3},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingPuzzle(tt.args.board); got != tt.want {
				t.Errorf("slidingPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}
