package week09

import (
	"testing"
)

func Test_shortestPathBinaryMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test05",
			args: args{
				grid: [][]int{{0}},
			},
			want: 1,
		},
		{
			name: "test04",
			args: args{
				grid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			},
			want: 4,
		},
		{
			name: "test03",
			args: args{
				grid: [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			},
			want: -1,
		},
		{
			name: "test02",
			args: args{
				grid: [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			},
			want: 4,
		},
		{
			name: "test01",
			args: args{
				grid: [][]int{
					{0, 1},
					{1, 0},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathBinaryMatrix(tt.args.grid); got != tt.want {
				t.Errorf("shortestPathBinaryMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortestPathBinaryMatrix2(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test06",
			args: args{
				grid: [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 1}},
			},
			want: -1,
		},
		{
			name: "test05",
			args: args{
				grid: [][]int{{0}},
			},
			want: 1,
		},
		{
			name: "test04",
			args: args{
				grid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			},
			want: 4,
		},
		{
			name: "test03",
			args: args{
				grid: [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			},
			want: -1,
		},
		{
			name: "test02",
			args: args{
				grid: [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			},
			want: 4,
		},
		{
			name: "test01",
			args: args{
				grid: [][]int{
					{0, 1},
					{1, 0},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathBinaryMatrix2(tt.args.grid); got != tt.want {
				t.Errorf("shortestPathBinaryMatrix2() = %v, want %v", got, tt.want)
			}
		})
	}
}
