package week01onclass

import "testing"

func Test_trap(t *testing.T) {
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
			name: "test01",
			args: args{
				heights: []int{0,1,0,2,1,0,1,3,2,1,2,1},
			},
			want: 6,
		},
		// {
		// 	name: "test02",
		// 	args: args{
		// 		heights: []int{4,2,0,3,2,5},
		// 	},
		// 	want: 9,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.heights); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
