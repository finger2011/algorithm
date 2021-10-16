package leetcode

import (
	"reflect"
	"testing"
)

func Test_addOperators(t *testing.T) {
	type args struct {
		num    string
		target int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		// {
		// 	name: "test01",
		// 	args: args{
		// 		num1:    "123",
		// 		target1: 6,
		// 	},
		// 	want: []string{"1+2+3","1*2*3"},
		// },
		// {
		// 	name: "test02",
		// 	args: args{
		// 		num:    "105",
		// 		target: 5,
		// 	},
		// 	want: []string{"1*0+5", "10-5"},
		// },
		{
			name: "test03",
			args: args{
				num:    "3456237490",
				target: 9191,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOperators(tt.args.num, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOperators() = %v, want %v", got, tt.want)
			}
		})
	}
}
