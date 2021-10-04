package week02onclass

import (
	"testing"
)

func Test_robotSim(t *testing.T) {
	type args struct {
		commands  []int
		obstacles [][]int
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
				commands:  []int{4, -1, 3},
				obstacles: [][]int{},
			},
			want: 25,
		},
		{
			name: "test02",
			args: args{
				commands:  []int{4, -1, 4, -2, 4},
				obstacles: [][]int{{2, 4}},
			},
			want: 65,
		},
		{
			name: "test03",
			args: args{
				commands:  []int{6, -1, -1, 6},
				obstacles: [][]int{},
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotSim(tt.args.commands, tt.args.obstacles); got != tt.want {
				t.Errorf("robotSim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_robotSim2(t *testing.T) {
	type args struct {
		commands  []int
		obstacles [][]int
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
				commands:  []int{4, -1, 3},
				obstacles: [][]int{},
			},
			want: 25,
		},
		{
			name: "test02",
			args: args{
				commands:  []int{4, -1, 4, -2, 4},
				obstacles: [][]int{{2, 4}},
			},
			want: 65,
		},
		{
			name: "test03",
			args: args{
				commands:  []int{6, -1, -1, 6},
				obstacles: [][]int{},
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotSim2(tt.args.commands, tt.args.obstacles); got != tt.want {
				t.Errorf("robotSim2() = %v, want %v", got, tt.want)
			}
		})
	}
}
