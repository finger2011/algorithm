package week05onclass

import "testing"

func Test_guessNumber(t *testing.T) {
	type args struct {
		n   int
		tar int
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
				n:   10,
				tar: 6,
			},
			want: 6,
		},
		{
			name: "test02",
			args: args{
				n:   1,
				tar: 1,
			},
			want: 1,
		},{
			name: "test03",
			args: args{
				n:   2,
				tar: 1,
			},
			want: 1,
		},
		{
			name: "test04",
			args: args{
				n:   2,
				tar: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target = tt.args.tar
			if got := guessNumber(tt.args.n); got != tt.want {
				t.Errorf("guessNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
