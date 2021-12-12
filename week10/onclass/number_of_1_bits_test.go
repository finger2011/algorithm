package week10onclass

import (
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
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
				num: 11,
			},
			want: 3,
		},
		{
			name: "test02",
			args: args{
				num: 128,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingWeight3(t *testing.T) {
	type args struct {
		num uint32
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
				num: 11,
			},
			want: 3,
		},
		{
			name: "test02",
			args: args{
				num: 128,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight3(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingWeight2(t *testing.T) {
	type args struct {
		num uint32
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
				num: 11,
			},
			want: 3,
		},
		{
			name: "test02",
			args: args{
				num: 128,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight2(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight2() = %v, want %v", got, tt.want)
			}
		})
	}
}
