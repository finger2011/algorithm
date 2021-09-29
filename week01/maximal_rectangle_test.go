package week01

import "testing"

//生成byte二维数组，与leetcode保持一致
func getBytes(a []string) [][]byte {
	// var a = []string{"10100", "10111", "11111", "10010"}
	var matrix = make([][]byte, len(a))
	for i := 0; i < len(a); i++ {
		matrix[i] = make([]byte, len(a[i]))
		for j := 0; j < len(a[i]); j++ {
			matrix[i][j] = a[i][j]
		}
	}
	return matrix
}

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
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
				matrix: getBytes([]string{"10100", "10111", "11111", "10010"}),
			},
			want: 6,
		},
		{
			name: "test02",
			args: args{
				matrix: [][]byte{},
			},
			want: 0,
		},
		{
			name: "test03",
			args: args{
				matrix: getBytes([]string{"0"}),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
