package week05

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{2, 2, 2, 0, 1},
			},
			want: 0,
		},
		{
			name: "test02",
			args: args{
				nums: []int{1, 3, 5},
			},
			want: 1,
		},
		{
			name: "test03",
			args: args{
				nums: []int{3, 1, 2, 3},
			},
			want: 1,
		},
		{
			name: "test04",
			args: args{
				nums: []int{3, 5, 6, 3},
			},
			want: 3,
		},
		{
			name: "test05",
			args: args{
				nums: []int{3, 3, 1, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
