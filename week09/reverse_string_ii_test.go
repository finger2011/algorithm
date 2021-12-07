package week09

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name:"test05",
			args:args{
				s:"abcdefg",
				k:8,
			},
			want:"gfedcba",
		},
		{
			name:"test04",
			args:args{
				s:"abcdefg",
				k:1,
			},
			want:"abcdefg",
		},
		{
			name:"test03",
			args:args{
				s:"a",
				k:2,
			},
			want:"a",
		},
		{
			name:"test01",
			args:args{
				s:"abcdefg",
				k:2,
			},
			want:"bacdfeg",
		},
		{
			name:"test02",
			args:args{
				s:"abcd",
				k:2,
			},
			want:"bacd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
