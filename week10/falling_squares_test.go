package week10

import (
	"reflect"
	"testing"
)

func Test_fallingSquares(t *testing.T) {
	type args struct {
		positions [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test03",
			args: args{
				positions: [][]int{{1, 5}, {2, 2}, {7, 5}},
			},
			want: []int{5, 7, 7},
		},
		{
			name: "test02",
			args: args{
				positions: [][]int{{9, 7}, {1, 9}, {3, 1}},
			},
			want: []int{7, 16, 17},
		},
		{
			name: "test01",
			args: args{
				positions: [][]int{{1, 2}, {2, 3}, {6, 1}},
			},
			want: []int{2, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fallingSquares(tt.args.positions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fallingSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
