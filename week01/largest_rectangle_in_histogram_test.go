package week01

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name : "test01",
			args : args{
				heights : []int{2,1,5,6,2,3},
			},
			want : 10,
		},
		{
			name : "test02",
			args : args{
				heights : []int{2,4},
			},
			want : 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("LargestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
