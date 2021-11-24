package week07onclass

import "testing"

func Test_rob2(t *testing.T) {
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
				nums: []int{2, 3, 2},
			},
			want: 3,
		},
		{
			name: "test02",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "test03",
			args: args{
				nums: []int{2},
			},
			want: 2,
		},
		{
			name: "test04",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob2(tt.args.nums); got != tt.want {
				t.Errorf("rob2s() = %v, want %v", got, tt.want)
			}
		})
	}
}
