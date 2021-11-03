package week05onclass

import "testing"

func Test_splitArray(t *testing.T) {
	type args struct {
		nums []int
		m    int
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
				nums: []int{7, 2, 5, 10, 8},
				m:    2,
			},
			want: 18,
		},
		{
			name: "test02",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				m:    2,
			},
			want: 9,
		},
		{
			name: "test01",
			args: args{
				nums: []int{1, 4, 4},
				m:    3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitArray(tt.args.nums, tt.args.m); got != tt.want {
				t.Errorf("splitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
