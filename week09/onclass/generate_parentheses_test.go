package week09onclass

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{n: 1},
			want: []string{"()"},
		},
		{
			name: "test02",
			args: args{n: 3},
			want: []string{"()()()", "()(())", "(())()", "(()())", "((()))"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !strEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func strEqual(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for _, str := range str1 {
		var isContain bool
		for _, compareStr := range str2 {
			if reflect.DeepEqual(str, compareStr) {
				isContain = true
				break
			}
		}
		if !isContain {
			return false
		}
	}
	return true
}
