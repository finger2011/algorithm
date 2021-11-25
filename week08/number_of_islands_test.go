package week08

import "testing"

func Test_numIslands(t *testing.T) {
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
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
