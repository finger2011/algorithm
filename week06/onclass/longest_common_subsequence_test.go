package week06onclasss

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
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
				text1: "abcde",
				text2: "ade",
			},
			want: 3,
		},
		{
			name: "test02",
			args: args{
				text1: "ade",
				text2: "ade",
			},
			want: 3,
		},
		{
			name: "test01",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}