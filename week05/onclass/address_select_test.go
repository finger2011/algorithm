package week05onclass

import "testing"

func Test_getMinDistance(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: []int{6, 2, 9, 1},
				k:    4,
			},
			want: 12,
		},
		{
			name: "test02",
			args: args{
				nums: []int{6, 2, 10},
				k:    3,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinDistance(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("getMinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
