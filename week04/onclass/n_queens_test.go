package week04onclass

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{n: 4},
			want: [][]string{
				[]string{".Q..", "...Q", "Q...", "..Q."},
				[]string{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			name: "test02",
			args: args{n: 1},
			want: [][]string{
				[]string{"Q"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
