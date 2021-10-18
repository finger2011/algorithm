package week04onclass

import (
	"testing"
)

func Test_numIslandsByBFS(t *testing.T) {
	type args struct {
		grid [][]byte
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
				grid: [][]byte{[]byte("11110"), []byte("11010"), []byte("11000"), []byte("00000")},
			},
			want: 1,
		},
		{
			name: "test02",
			args: args{
				grid: [][]byte{[]byte("11000"), []byte("11000"), []byte("00100"), []byte("00011")},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslandsByBFS(tt.args.grid); got != tt.want {
				t.Errorf("numIslandsByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numIslandsByDFS(t *testing.T) {
	type args struct {
		grid [][]byte
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
				grid: [][]byte{[]byte("11110"), []byte("11010"), []byte("11000"), []byte("00000")},
			},
			want: 1,
		},
		{
			name: "test02",
			args: args{
				grid: [][]byte{[]byte("11000"), []byte("11000"), []byte("00100"), []byte("00011")},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslandsByDFS(tt.args.grid); got != tt.want {
				t.Errorf("numIslandsByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
