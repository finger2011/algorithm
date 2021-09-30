package week01onclass

import "testing"

func Test_calculateWithoutParenthess(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{s: "3+2*2"},
			want: 7,
		},
		{
			name: "test02",
			args: args{s: "3/2"},
			want: 1,
		},
		{
			name: "test03",
			args: args{s: " 3+5 / 2 "},
			want: 5,
		},
		{
			name: "test04",
			args: args{s: "1*2-3/4+5*6-7*8+9/10"},
			want: -24,
		},
		{
			name: "test05",
			args: args{s: "1-1+1"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWithoutParenthess(tt.args.s); got != tt.want {
				t.Errorf("calculateWithoutParenthess() = %v, want %v", got, tt.want)
			}
		})
	}
}
