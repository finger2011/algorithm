package week09

import "testing"

func Test_reverseWords3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				s: "Let's take LeetCode contest",
			},
			want: "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords3(tt.args.s); got != tt.want {
				t.Errorf("reverseWords3() = %v, want %v", got, tt.want)
			}
		})
	}
}
